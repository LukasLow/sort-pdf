<script>
    import { getFilename } from "../../lib/formState.svelte.js";
</script>

<div class="p-4 rounded-2xl bg-card border border-border shadow-2xl">
    <div class="text-sm font-semibold text-primary mb-2">Vorschau Dateiname</div>
    <div class="text-base font-mono font-semibold text-primary break-all">
        {getFilename()}
    </div>
</div>
