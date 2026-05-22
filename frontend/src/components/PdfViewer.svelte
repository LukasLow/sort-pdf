<script>
  import { onMount } from 'svelte';
  import * as pdfjsLib from 'pdfjs-dist';
  import pdfWorker from 'pdfjs-dist/build/pdf.worker.min.mjs?url';

  pdfjsLib.GlobalWorkerOptions.workerSrc = pdfWorker;

  let { filename } = $props();

  let container;
  let pdfDoc = null;

  let zoom = $state(1);
  let loading = $state(false);

  async function loadPdf() {
    if (!filename) return;

    loading = true;

    try {
      const url = `http://127.0.0.1:34999/pdf/${filename}`;

      pdfDoc = await pdfjsLib
        .getDocument(url)
        .promise;

      await renderPdf();
    } catch (err) {
      console.error('PDF load error:', err);
    } finally {
      loading = false;
    }
  }

  async function renderPdf() {
    if (!pdfDoc || !container) return;

    container.innerHTML = '';

    const containerWidth =
      container.clientWidth || 800;

    for (let i = 1; i <= pdfDoc.numPages; i++) {
      const page = await pdfDoc.getPage(i);

      const baseViewport =
        page.getViewport({ scale: 1 });

      // Seitenbreite automatisch anpassen
      const fitScale =
        containerWidth / baseViewport.width;

      const viewport = page.getViewport({
        scale: fitScale * zoom
      });

      const canvas =
        document.createElement('canvas');

      const ctx = canvas.getContext('2d');

      canvas.width = viewport.width;
      canvas.height = viewport.height;

      canvas.className =
        'mb-6 rounded shadow-xl bg-white w-full';

      await page.render({
        canvasContext: ctx,
        viewport
      }).promise;

      container.appendChild(canvas);
    }
  }

  function handleWheel(e) {
    // Nur zoomen wenn STRG/CMD gedrückt
    // dadurch normales Scrollen wieder möglich
    if (!e.ctrlKey && !e.metaKey) {
      return;
    }

    e.preventDefault();

    if (e.deltaY < 0) {
      zoom = Math.min(zoom + 0.1, 3);
    } else {
      zoom = Math.max(zoom - 0.1, 0.5);
    }

    renderPdf();
  }

  onMount(() => {
    loadPdf();

    // Resize support
    window.addEventListener('resize', renderPdf);
  });

  $effect(() => {
    if (filename) {
      loadPdf();
    }
  });
</script>

<div
  onwheel={handleWheel}
  class="w-2/3 h-full bg-zinc-900 border-r border-zinc-800 flex flex-col relative overflow-hidden"
>
  {#if filename}
    <div class="flex-1 overflow-auto">
      <div
        bind:this={container}
        class="w-full p-6 flex flex-col"
      ></div>
    </div>

    <div
      class="absolute top-3 right-3 bg-zinc-950/80 backdrop-blur border border-zinc-800 px-2.5 py-1 rounded-md text-[10px] font-mono text-zinc-400"
    >
      Zoom: {Math.round(zoom * 100)}%
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
