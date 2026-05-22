<script lang="ts">
  // Svelte 5 Props
  let { workDir, onSelect }: { workDir: string, onSelect: () => void } = $props();
</script>

<div class="w-1/3 h-full p-6 flex flex-col justify-between bg-zinc-950 overflow-y-auto border-l border-zinc-900">
  <div class="space-y-6">
    <div class="space-y-1">
      <h2 class="text-lg font-bold text-white tracking-tight">Dokumenten-Aktionen</h2>
      <p class="text-xs text-zinc-400">PDF einsortieren und verschlagworten</p>
    </div>

    <div class="h-px bg-zinc-800 w-full"></div>

    <!-- Under Construction Platzhalter -->
    <div class="flex flex-col items-center justify-center border-2 border-dashed border-zinc-800 rounded-2xl p-8 text-center space-y-3 my-8">
      <span class="text-4xl animate-bounce">🚧</span>
      <div class="space-y-1">
        <h3 class="text-sm font-semibold text-zinc-200">Seitenliste & Formular</h3>
        <p class="text-xs text-zinc-500 max-w-[200px]">Hier entstehen bald die Dropdowns für Korrespondenten und das Datum.</p>
      </div>
    </div>
  </div>

  <!-- Unterer Bereich: Arbeitsordner-Kontrolle & Version -->
  <div class="space-y-4 pt-4 border-t border-zinc-900">
    <div class="bg-zinc-900/50 border border-zinc-800 rounded-xl p-3 break-all space-y-2">
      <div>
        <div class="text-[9px] font-bold text-blue-400 uppercase tracking-wider mb-0.5">Aktiver Pfad</div>
        <div class="text-[11px] font-mono text-zinc-400 select-all">{workDir}</div>
      </div>

      <button
        onclick={onSelect}
        class="w-full bg-zinc-800 hover:bg-zinc-700 active:scale-[0.98] transition-transform text-zinc-300 rounded-lg py-1.5 text-[11px] font-medium border border-zinc-700/30"
      >
        📁 Ordner ändern
      </button>
    </div>

    <div class="text-[10px] text-zinc-600 font-mono text-center">
      sort-pdf v0.1.0
    </div>
  </div>
</div>
