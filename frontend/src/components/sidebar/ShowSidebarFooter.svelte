<script>
    let { workDir, onSelect } = $props();
</script>

<div
    class="mt-auto p-4 rounded-2xl bg-zinc-900 border border-zinc-700 rounded-2xl shadow-2xl"
>
    <div class="text-sm font-semibold text-blue-400 mb-2">
        Arbeitsverzeichnis
    </div>
    <div class="text-base font-mono font-semibold text-blue-300 break-all">
        {workDir}
    </div>
    <button
        on:click={onSelect}
        class="mt-2 w-full bg-green-600 text-blue-300 px-4 py-2 rounded-lg text-sm font-semibold hover:bg-green-700"
    >
        Verzeichnis auswählen
    </button>
</div>
