<script lang="ts">
  let { onSelect }: { onSelect: () => void } = $props();
</script>

<main class="flex h-screen w-screen items-center justify-center bg-background text-foreground antialiased p-8">
  <div class="max-w-md w-full bg-card border border-border rounded-2xl p-6 shadow-2xl space-y-6">
    <div class="space-y-1">
      <h1 class="text-xl font-bold tracking-tight text-foreground">PDF-Sortierer</h1>
      <p class="text-xs text-muted-foreground">Arbeitsumgebung einrichten</p>
    </div>
    <div class="h-px bg-border-dim w-full"></div>
    <p class="text-sm text-muted-foreground leading-relaxed">
      Es wurde noch kein Arbeitsverzeichnis konfiguriert. Bitte wähle einen Ordner aus.
    </p>
    <button
      onclick={onSelect}
      class="w-full bg-primary hover:opacity-85 active:scale-[0.98] transition-all text-background rounded-xl py-3 font-semibold text-sm shadow-lg shadow-primary/10"
    >
      Ordner auswählen
    </button>
  </div>
</main>
