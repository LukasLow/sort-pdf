<script>
  // Svelte 5 Props in reinem JavaScript mit Two-Way-Binding für currentPdf
  let { workDir, pdfList, correspondents, currentPdf = $bindable(), onSelect } = $props();
</script>

<div class="w-1/3 h-full p-6 flex flex-col justify-between bg-zinc-950 overflow-y-auto border-l border-zinc-900">
  <div class="space-y-6">
    <div class="space-y-1">
      <h2 class="text-lg font-bold text-white tracking-tight">Dokumenten-Aktionen</h2>
      <p class="text-xs text-zinc-400">PDF einsortieren und verschlagworten</p>
    </div>

    <div class="h-px bg-zinc-800 w-full"></div>

    <!-- Sektion 1: Liste der PDFs im Eingang -->
    <div class="space-y-2">
      <label class="text-xs font-semibold text-zinc-400 uppercase tracking-wider">Eingang ({pdfList.length})</label>
      {#if pdfList.length === 0}
        <div class="text-xs text-zinc-600 bg-zinc-900/30 p-3 rounded-xl border border-zinc-900 border-dashed">
          Keine PDFs im Ordner 900-Eingang gefunden.
        </div>
      {:else}
        <div class="max-h-40 overflow-y-auto space-y-1 pr-1 border border-zinc-900 p-2 rounded-xl bg-zinc-900/20">
          {#each pdfList as pdf}
            <button
              onclick={() => currentPdf = pdf}
              class="w-full text-left px-3 py-2 rounded-lg text-xs font-mono truncate transition-colors {currentPdf === pdf ? 'bg-blue-600 text-white font-semibold shadow-md' : 'text-zinc-400 hover:bg-zinc-900 hover:text-zinc-200'}"
            >
              📄 {pdf}
            </button>
          {/each}
        </div>
      {/if}
    </div>

    <!-- Sektion 2: Das Sortier-Formular -->
    <div class="space-y-4 pt-2">
      <label class="text-xs font-semibold text-zinc-400 uppercase tracking-wider block">Klassifizierung</label>
      
      <div class="space-y-1.5">
        <span class="text-xs text-zinc-500">Korrespondent (aus YAML):</span>
        <select class="w-full bg-zinc-900 border border-zinc-800 rounded-xl p-2.5 text-xs text-zinc-200 focus:outline-none focus:border-blue-500">
          <option value="">-- Bitte wählen --</option>
          {#each correspondents as corp}
            <option value={corp}>{corp}</option>
          {/each}
        </select>
      </div>

      <!-- Kleines Info-Feld -->
      <div class="flex items-center gap-2 p-3 bg-zinc-900/40 rounded-xl border border-zinc-800/60 text-xs text-zinc-500">
        <span>🚧</span>
        <span>Datumserkennung & OCR folgen in Phase 4.</span>
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