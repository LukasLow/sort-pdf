<script>
    import { form } from "../../lib/formState.svelte.js";

    const years = Array.from({ length: 126 }, (_, i) => 2026 - i);
    const months = Array.from({ length: 12 }, (_, i) => i + 1);
    const days = Array.from({ length: 31 }, (_, i) => i + 1);
</script>

<div
    class="p-4 rounded-2xl bg-zinc-900 border border-zinc-700 shadow-2xl space-y-2"
>
    <div class="text-sm font-semibold text-blue-400">Datum (für Filename)</div>

    <div class="flex gap-2">
        <select
            bind:value={form.year}
            class="bg-zinc-800 border border-zinc-600 text-base font-semibold text-blue-300 rounded-lg p-2 text-xs w-full shadow-[inset_2px_2px_5px_rgba(0,0,0,0.5)]"
        >
            {#each years as y}
                <option value={y}>{y}</option>
            {/each}
        </select>

        <select
            bind:value={form.month}
            class="bg-zinc-800 border border-zinc-600 text-base font-semibold text-blue-300 rounded-lg p-2 text-xs w-full shadow-[inset_2px_2px_5px_rgba(0,0,0,0.5)]"
        >
            {#each months as m}
                <option value={m}>{m.toString().padStart(2, "0")}</option>
            {/each}
        </select>

        <select
            bind:value={form.day}
            class="bg-zinc-800 border border-zinc-600 text-base font-semibold text-blue-300 rounded-lg p-2 text-xs w-full shadow-[inset_2px_2px_5px_rgba(0,0,0,0.5)]"
        >
            {#each days as d}
                <option value={d}>{d.toString().padStart(2, "0")}</option>
            {/each}
        </select>
    </div>

    <div class="text-xs font-semibold text-blue-400">Format: YYYY-MM-DD</div>
</div>
