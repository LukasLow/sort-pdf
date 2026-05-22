<script>
  let { currentPdf } = $props();
</script>

<div class="p-4 rounded-2xl bg-white/5 border border-white/10 backdrop-blur-xl shadow-lg">
  <div class="text-xs text-zinc-500 mb-2">Aktuelle Datei</div>

  {#if currentPdf}
    <div class="text-sm font-mono text-zinc-200 break-all">
      {currentPdf}
    </div>
  {:else}
    <div class="text-xs text-zinc-600">
      Keine PDF ausgewählt
    </div>
  {/if}

  <div class="text-[10px] text-zinc-600 mt-2">
    🚧 später: Preview + OCR + Analyse
  </div>
</div>
