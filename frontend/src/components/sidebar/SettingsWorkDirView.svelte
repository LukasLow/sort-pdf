<script>
    let { workDir, onSelect } = $props();
</script>

<div>
    <div class="text-lg font-bold text-primary mb-4">Arbeitsverzeichnis</div>

    <div class="p-4 rounded-xl bg-muted border border-border-dim">
        <div class="text-sm font-semibold text-muted-foreground mb-2">
            Aktuelles Verzeichnis
        </div>
        <div class="text-base font-mono text-primary break-all mb-4">
            {workDir || "Kein Verzeichnis ausgewählt"}
        </div>
        <button
            onclick={onSelect}
            class="w-full bg-primary text-background px-4 py-2 rounded-lg text-sm font-semibold hover:bg-primary-dark"
        >
            Verzeichnis auswählen
        </button>
    </div>
</div>
