<script>
    let { previewUrl } = $props();
</script>

<div
    class="p-4 rounded-2xl bg-zinc-900 border border-zinc-700 rounded-2xl shadow-2xl"
>
    <div class="text-sm font-semibold text-blue-400 mb-2">Vorschau</div>
    {#if previewUrl}
        <img src={previewUrl} alt="Preview" class="w-full h-auto rounded-lg" />
    {:else}
        <div class="text-sm font-semibold text-blue-400">
            Keine Vorschau verfügbar
        </div>
    {/if}
</div>
