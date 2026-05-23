<script>
    import { form } from "../../lib/formState.svelte.js";
</script>

<div
    class="p-4 rounded-2xl bg-zinc-900 border border-zinc-700 shadow-2xl space-y-2"
>
    <div class="text-sm font-semibold text-blue-400">Info (Filename Teil)</div>

    <input
        bind:value={form.info}
        class="w-full bg-zinc-800 border border-zinc-600 text-base font-semibold text-blue-300 rounded-lg p-2 text-xs text-zinc-200 shadow-[inset_2px_2px_5px_rgba(0,0,0,0.5)]"
        placeholder="z.B. Rechnung_Miete"
    />

    <div class="text-xs font-semibold text-blue-400">
        YYYY-MM-DD_KORRESPONDENT_INFO_EXTRAS.pdf
    </div>
</div>
