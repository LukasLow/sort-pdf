<script>
    import { form } from "../../lib/formState.svelte.js";
</script>

<div
    class="p-4 rounded-2xl bg-card border border-border shadow-2xl space-y-2"
>
    <div class="text-sm font-semibold text-primary">Extras</div>

    <input
        bind:value={form.extras}
        class="w-full bg-muted border border-border-dim text-base font-semibold text-primary rounded-lg p-2 text-xs text-foreground shadow-[inset_2px_2px_5px_rgba(0,0,0,0.5)]"
        placeholder="z.B. Rechnungsnummer"
    />

    <div class="text-xs font-semibold text-primary">
        YYYY-MM-DD_KORRESPONDENT_INFO_EXTRAS.pdf
    </div>
</div>
