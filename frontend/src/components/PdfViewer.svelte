<script>
    import { onMount, onDestroy } from "svelte";
    import { GetPDFPageCount } from "../../wailsjs/go/src/WorkspaceBridge";

    let { filename } = $props();

    let loading = $state(false);
    let scrollEl = $state();
    let zoom = $state(1);
    let containerWidth = $state(800);
    let pageCount = $state(0);
    let resizeTimer;

    function pdfUrl() {
        if (!filename) return "";
        return `http://127.0.0.1:34999/pdf/${encodeURIComponent(filename)}`;
    }

    function applyZoom(newZoom) {
        newZoom = Math.max(0.5, Math.min(3, newZoom));
        if (newZoom === zoom) return;
        const oldZoom = zoom;
        zoom = newZoom;
        if (!scrollEl) return;
        const ratio = newZoom / oldZoom;
        const el = scrollEl;
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

    function handleKeyDown(e) {
        if (!e.metaKey && !e.ctrlKey) return;
        if (e.target.closest("input,textarea,select")) return;
        if (e.key === "=" || e.key === "+") { e.preventDefault(); zoomIn(); }
        else if (e.key === "-") { e.preventDefault(); zoomOut(); }
        else if (e.key === "0") { e.preventDefault(); zoomReset(); }
    }

    function updateContainerWidth() {
        if (scrollEl) containerWidth = scrollEl.clientWidth;
    }

    function handleResize() {
        clearTimeout(resizeTimer);
        resizeTimer = setTimeout(() => updateContainerWidth(), 150);
    }

    onMount(() => {
        updateContainerWidth();
        window.addEventListener("resize", handleResize);
        window.addEventListener("keydown", handleKeyDown);
    });

    onDestroy(() => {
        window.removeEventListener("resize", handleResize);
        window.removeEventListener("keydown", handleKeyDown);
        clearTimeout(resizeTimer);
    });

    $effect(() => {
        if (filename) {
            loading = true;
            pageCount = 0;
            updateContainerWidth();
            getPageCount()
                .then(c => pageCount = c)
                .finally(() => loading = false);
        }
    });

    async function getPageCount() {
        if (!filename) return 0;
        try {
            return await GetPDFPageCount(filename);
        } catch {
            return 0;
        }
    }
</script>

<div class="flex-1 h-full bg-zinc-900 border-l border-zinc-800 flex flex-col relative overflow-hidden">
    {#if filename}
        <div bind:this={scrollEl} class="flex-1 overflow-auto">
            <div class="p-6 flex flex-col gap-6" style="width: fit-content;">
                <div style="transform-origin: top left; transform: scale({zoom});">
                    <embed
                        src={pdfUrl()}
                        type="application/pdf"
                        class="block"
                        style="width: {containerWidth}px; height: {containerWidth * 1.4}px; min-height: 400px;"
                    />
                </div>
            </div>
        </div>

        <div class="absolute top-3 right-3 flex items-center gap-1 bg-zinc-950/80 backdrop-blur border border-zinc-800 px-2 py-1 rounded-md text-xs font-mono text-zinc-400 select-none z-10">
            <button onclick={zoomOut} class="hover:text-zinc-200 px-1 leading-none text-sm">−</button>
            <span class="min-w-[4ch] text-center">{Math.round(zoom * 100)}%</span>
            <button onclick={zoomIn} class="hover:text-zinc-200 px-1 leading-none text-sm">+</button>
            {#if pageCount > 0}
                <span class="ml-2 text-zinc-500">| {pageCount} {pageCount === 1 ? "Seite" : "Seiten"}</span>
            {/if}
        </div>

        {#if loading}
            <div class="absolute inset-0 flex items-center justify-center bg-zinc-900/70 text-zinc-400 text-sm pointer-events-none z-20">
                PDF wird geladen...
            </div>
        {/if}
    {:else}
        <div class="flex h-full w-full items-center justify-center text-zinc-500 text-sm">
            Keine PDF geladen
        </div>
    {/if}
</div>
