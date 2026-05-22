<script lang="ts">
  import { onMount } from 'svelte';
  // Die von Wails frisch generierten Bindings für unsere Brücke:
  import { CheckInitialWorkDir, SelectWorkingDirectory } from '../wailsjs/go/src/WorkspaceBridge';

  // Svelte 5 Zustands-Runes
  let workDir = $state('');
  let errorMessage = $state('');

  // Sobald die App lädt, prüfen wir, ob macOS ein gespeichertes Workdir hat
  onMount(async () => {
    try {
      workDir = await CheckInitialWorkDir();
    } catch (err) {
      errorMessage = 'Fehler beim Laden der Konfiguration';
    }
  });

  // Öffnet den nativen macOS-Dialog über unsere Go-Brücke
  async function handleFolderSelection() {
    try {
      errorMessage = '';
      const selected = await SelectWorkingDirectory();
      if (selected) {
        workDir = selected;
      }
    } catch (err) {
      errorMessage = String(err);
    }
  }
</script>

<main class="flex h-screen w-screen items-center justify-center bg-zinc-950 text-zinc-100 antialiased p-8 font-sans">
  
  <div class="max-w-md w-full bg-zinc-900 border border-zinc-800 rounded-2xl p-6 shadow-2xl space-y-6">
    
    <div class="space-y-1">
      <h1 class="text-xl font-bold tracking-tight text-white">PDF-Sortierer</h1>
      <p class="text-xs text-zinc-400">Session-übergreifende Arbeitsumgebung</p>
    </div>

    <div class="h-px bg-zinc-800 w-full"></div>

    {#if !workDir}
      <!-- Zustand: Kein Ordner hinterlegt -->
      <div class="space-y-4">
        <p class="text-sm text-zinc-400 leading-relaxed">
          Es wurde noch kein Arbeitsverzeichnis konfiguriert. Bitte wähle einen Ordner aus, in dem die Archiv-Struktur angelegt werden soll.
        </p>
        
        <button 
          onclick={handleFolderSelection} 
          class="w-full bg-blue-600 hover:bg-blue-500 active:scale-[0.98] transition-transform text-white rounded-xl py-3 font-semibold text-sm shadow-lg shadow-blue-600/10"
        >
          📁 Ordner auswählen
        </button>
      </div>
    {:else}
      <!-- Zustand: Ordner erfolgreich geladen/ausgewählt -->
      <div class="space-y-4">
        <div class="bg-zinc-950 border border-zinc-800 rounded-xl p-3.5 break-all">
          <div class="text-[10px] font-bold text-blue-400 uppercase tracking-wider mb-1">Aktiver Pfad</div>
          <div class="text-xs font-mono text-zinc-300">{workDir}</div>
        </div>

        <p class="text-xs text-emerald-400 flex items-center gap-1.5">
          ✓ Struktur ist bereit und verknüpft.
        </p>

        <button 
          onclick={handleFolderSelection} 
          class="w-full bg-zinc-800 hover:bg-zinc-700 text-zinc-200 rounded-xl py-2 text-xs font-medium border border-zinc-700/50"
        >
          Ordner ändern
        </button>
      </div>
    {/if}

    {#if errorMessage}
      <!-- Fehleranzeige, falls ein Schritt fehlschlägt -->
      <div class="text-xs bg-red-950/50 border border-red-900/50 text-red-400 p-3 rounded-lg font-mono">
        {errorMessage}
      </div>
    {/if}

  </div>

</main>