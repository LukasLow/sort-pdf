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
        } catch (err) {
            console.error("PDF load error:", err);
        } finally {
            loading = false;
        }
    }

    function applyZoom(newZoom) {
        const oldZoom = zoom;
        newZoom = Math.max(0.5, Math.min(3, newZoom));
        if (newZoom === oldZoom) return;
        zoom = newZoom;
        if (!scrollContainer) return;
        const ratio = newZoom / oldZoom;
        const el = scrollContainer;
        requestAnimationFrame(() => {
            const cx = el.clientWidth / 2;
            const cy = el.clientHeight / 2;
            el.scrollLeft = el.scrollLeft * ratio + cx * (ratio - 1);
            el.scrollTop = el.scrollTop * ratio + cy * (ratio - 1);
        });
    }

    function zoomIn() { applyZoom(zoom + 0.1); }
    function zoomOut() { applyZoom(zoom - 0.1); }
    function zoomReset() { applyZoom(1); }

    function handleWheel(e) {
        if (!e.ctrlKey && !e.metaKey) return;
        e.preventDefault();
        const oldZoom = zoom;
        const newZoom = Math.max(0.5, Math.min(3, oldZoom + (e.deltaY < 0 ? 0.1 : -0.1)));
        if (newZoom === oldZoom) return;
        const rect = scrollContainer.getBoundingClientRect();
        const mx = e.clientX - rect.left;
        const my = e.clientY - rect.top;
        const sl = scrollContainer.scrollLeft;
        const st = scrollContainer.scrollTop;
        zoom = newZoom;
        const ratio = newZoom / oldZoom;
        requestAnimationFrame(() => {
            scrollContainer.scrollLeft = sl * ratio + mx * (ratio - 1);
            scrollContainer.scrollTop = st * ratio + my * (ratio - 1);
        });
    }

    function handleKeyDown(e) {
        if (!e.metaKey && !e.ctrlKey) return;
        if (e.target.closest("input,textarea,select")) return;
        if (e.key === "=" || e.key === "+") { e.preventDefault(); zoomIn(); }
        else if (e.key === "-") { e.preventDefault(); zoomOut(); }
        else if (e.key === "0") { e.preventDefault(); zoomReset(); }
    }

    let renderTick = $state(0);

    function handleResize() {
        clearTimeout(resizeTimer);
        resizeTimer = setTimeout(() => { renderTick++; }, 150);
    }

    onMount(() => {
        loadPdf();
        window.addEventListener("resize", handleResize);
        window.addEventListener("keydown", handleKeyDown);
    });

    onDestroy(() => {
        window.removeEventListener("resize", handleResize);
        window.removeEventListener("keydown", handleKeyDown);
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
            onwheel={handleWheel}
        >
            <div class="w-full p-6 flex flex-col gap-6">
                {#each pages as pageNum}
                    <PageRenderer {pdfDoc} {pageNum} {zoom} tick={renderTick} />
                {/each}
            </div>
        </div>

        <div class="absolute top-3 right-3 flex items-center gap-1 bg-zinc-950/80 backdrop-blur border border-zinc-800 px-2 py-1 rounded-md text-xs font-mono text-zinc-400 select-none">
            <button onclick={zoomOut} class="hover:text-zinc-200 px-1 leading-none text-sm">−</button>
            <span class="min-w-[4ch] text-center">{Math.round(zoom * 100)}%</span>
            <button onclick={zoomIn} class="hover:text-zinc-200 px-1 leading-none text-sm">+</button>
        </div>

        {#if loading}
            <div class="absolute inset-0 flex items-center justify-center bg-zinc-900/70 text-zinc-400 text-sm pointer-events-none">
                PDF wird geladen...
            </div>
        {/if}
    {:else}
        <div class="flex h-full w-full items-center justify-center text-zinc-500 text-sm">
            Keine PDF geladen
        </div>
    {/if}
</div>
