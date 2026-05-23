<script>
    let { workDir, onSelect } = $props();
</script>

<div>
    <div class="text-lg font-bold text-blue-300 mb-4">Arbeitsverzeichnis</div>

    <div class="p-4 rounded-xl bg-zinc-800 border border-zinc-700">
        <div class="text-sm font-semibold text-zinc-400 mb-2">
            Aktuelles Verzeichnis
        </div>
        <div class="text-base font-mono text-blue-300 break-all mb-4">
            {workDir || "Kein Verzeichnis ausgewählt"}
        </div>
        <button
            onclick={onSelect}
            class="w-full bg-green-600 text-white px-4 py-2 rounded-lg text-sm font-semibold hover:bg-green-700"
        >
            Verzeichnis auswählen
        </button>
    </div>
</div>
