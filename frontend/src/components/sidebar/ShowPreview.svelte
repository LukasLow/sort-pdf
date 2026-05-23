<script>
    import { getFilename } from "../../lib/formState.svelte.js";
</script>

<div class="p-4 rounded-2xl bg-zinc-900 border border-zinc-700 shadow-2xl">
    <div class="text-sm font-semibold text-blue-400 mb-2">Vorschau Dateiname</div>
    <div class="text-base font-mono font-semibold text-blue-300 break-all">
        {getFilename()}
    </div>
</div>
