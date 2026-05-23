<script>
    import { onMount, onDestroy } from "svelte";
    import { getDocument, GlobalWorkerOptions } from "pdfjs-dist";
    import workerUrl from "pdfjs-dist/build/pdf.worker.min.mjs?url";
    import PageRenderer from "./PageRenderer.svelte";

    GlobalWorkerOptions.workerSrc = workerUrl;

    let { filename } = $props();

    let pdfDoc = $state(null);
    let pageCount = $state(0);
    let zoom = $state(1);
    let loading = $state(false);
    let renderedUpTo = $state(0);
    let pages = $state([]);
    let resizeTimer;
    let scrollContainer = $state();

    async function loadPdf() {
        if (!filename) return;
        loading = true;
        pdfDoc = null;
        pages = [];

        try {
            const url = `http://127.0.0.1:34999/pdf/${filename}`;
            pdfDoc = await getDocument(url).promise;
            pageCount = pdfDoc.numPages;
            pages = Array.from({ length: pageCount }, (_, i) => i + 1);
            renderedUpTo = Math.min(3, pageCount);
        } catch (err) {
            console.error("PDF load error:", err);
        } finally {
            loading = false;
        }
    }

    function handleScroll(e) {
        if (renderedUpTo >= pageCount) return;
        const el = e.target;
        if (el.scrollTop + el.clientHeight >= el.scrollHeight - 600) {
            renderedUpTo = Math.min(renderedUpTo + 2, pageCount);
        }
    }

    function handleWheel(e) {
        if (!e.ctrlKey && !e.metaKey) return;
        e.preventDefault();

        const oldZoom = zoom;
        const newZoom = Math.max(0.5, Math.min(3, oldZoom + (e.deltaY < 0 ? 0.1 : -0.1)));
        if (newZoom === oldZoom) return;

        const rect = scrollContainer.getBoundingClientRect();
        const mx = e.clientX - rect.left;
        const my = e.clientY - rect.top;
        const scrollLeft = scrollContainer.scrollLeft;
        const scrollTop = scrollContainer.scrollTop;

        zoom = newZoom;
        const ratio = newZoom / oldZoom;

        requestAnimationFrame(() => {
            scrollContainer.scrollLeft = scrollLeft * ratio + mx * (ratio - 1);
            scrollContainer.scrollTop = scrollTop * ratio + my * (ratio - 1);
        });
    }

    let renderTick = $state(0);

    function handleResize() {
        clearTimeout(resizeTimer);
        resizeTimer = setTimeout(() => {
            renderTick++;
        }, 150);
    }

    onMount(() => {
        loadPdf();
        window.addEventListener("resize", handleResize);
    });

    onDestroy(() => {
        window.removeEventListener("resize", handleResize);
        clearTimeout(resizeTimer);
    });

    $effect(() => {
        if (filename) loadPdf();
    });
</script>

<div
    class="w-full h-full bg-zinc-900 border-l border-zinc-800 flex flex-col relative overflow-hidden"
>
    {#if filename}
        <div
            bind:this={scrollContainer}
            class="flex-1 overflow-auto"
            onscroll={handleScroll}
            onwheel={handleWheel}
        >
            <div class="w-full p-6 flex flex-col gap-6">
                {#each pages as pageNum}
                    {#if pageNum <= renderedUpTo}
                        <PageRenderer
                            {pdfDoc}
                            {pageNum}
                            {zoom}
                            tick={renderTick}
                        />
                    {/if}
                {/each}
                {#if renderedUpTo < pageCount}
                    <div class="flex justify-center py-4">
                        <span class="text-xs text-zinc-500">Scrolle fuer weitere Seiten...</span>
                    </div>
                {/if}
            </div>
        </div>

        <div
            class="absolute top-3 right-3 bg-zinc-950/80 backdrop-blur border border-zinc-800 px-2.5 py-1 rounded-md text-[10px] font-mono text-zinc-400"
        >
            Zoom: {Math.round(zoom * 100)}% - Seite 1-{renderedUpTo} von {pageCount}
        </div>

        {#if loading}
            <div
                class="absolute inset-0 flex items-center justify-center bg-zinc-900/70 text-zinc-400 text-sm"
            >
                PDF wird geladen...
            </div>
        {/if}
    {:else}
        <div
            class="flex h-full w-full items-center justify-center text-zinc-500 text-sm"
        >
            Keine PDF geladen
        </div>
    {/if}
</div>

