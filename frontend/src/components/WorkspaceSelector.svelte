<script lang="ts">
  let { onSelect }: { onSelect: () => void } = $props();
</script>

<main class="flex h-screen w-screen items-center justify-center bg-zinc-950 text-zinc-100 antialiased p-8">
  <div class="max-w-md w-full bg-zinc-900 border border-zinc-800 rounded-2xl p-6 shadow-2xl space-y-6">
    <div class="space-y-1">
      <h1 class="text-xl font-bold tracking-tight text-white">PDF-Sortierer</h1>
      <p class="text-xs text-zinc-400">Arbeitsumgebung einrichten</p>
    </div>
    <div class="h-px bg-zinc-800 w-full"></div>
    <p class="text-sm text-zinc-400 leading-relaxed">
      Es wurde noch kein Arbeitsverzeichnis konfiguriert. Bitte wähle einen Ordner aus.
    </p>
    <button
      onclick={onSelect}
      class="w-full bg-blue-600 hover:bg-blue-500 active:scale-[0.98] transition-transform text-white rounded-xl py-3 font-semibold text-sm shadow-lg shadow-blue-600/10"
    >
      📁 Ordner auswählen
    </button>
  </div>
</main>
