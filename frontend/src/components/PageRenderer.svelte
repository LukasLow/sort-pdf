<script>
    import { TextLayer } from "pdfjs-dist";

    let { pdfDoc, pageNum, zoom, tick = 0 } = $props();

    let container = $state();
    let canvas = $state();
    let textLayerDiv = $state();
    let curTextLayer = null;
    let curRenderTask = null;

    $effect(() => {
        tick;
        if (canvas && pdfDoc && container && textLayerDiv) render();
    });

    async function render() {
        if (!pdfDoc || !canvas || !container || !textLayerDiv) return;

        if (curRenderTask) {
            try { curRenderTask.cancel(); } catch {}
            curRenderTask = null;
        }
        if (curTextLayer) {
            curTextLayer.cancel();
            curTextLayer = null;
        }

        const page = await pdfDoc.getPage(pageNum);
        const containerWidth = container.clientWidth || 800;
        const base = page.getViewport({ scale: 1 });
        const scale = (containerWidth / base.width) * zoom;
        const viewport = page.getViewport({ scale });

        const w = Math.floor(viewport.width);
        const h = Math.floor(viewport.height);

        canvas.width = w * window.devicePixelRatio;
        canvas.height = h * window.devicePixelRatio;
        canvas.style.width = w + "px";
        canvas.style.height = h + "px";

        const ctx = canvas.getContext("2d");
        ctx.scale(window.devicePixelRatio, window.devicePixelRatio);

        curRenderTask = page.render({ canvasContext: ctx, viewport });
        await curRenderTask.promise;
        curRenderTask = null;

        const textContent = await page.getTextContent();

        textLayerDiv.innerHTML = "";
        textLayerDiv.style.width = w + "px";
        textLayerDiv.style.height = h + "px";

        curTextLayer = new TextLayer({
            textContentSource: textContent,
            container: textLayerDiv,
            viewport,
        });
        await curTextLayer.render();
    }
</script>

<div bind:this={container} class="w-full relative">
    <canvas bind:this={canvas} class="rounded shadow-xl bg-white block"></canvas>
    <div
        bind:this={textLayerDiv}
        class="absolute top-0 left-0 text-layer"
    ></div>
</div>
