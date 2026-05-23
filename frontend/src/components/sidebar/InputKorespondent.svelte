<script>
    import { form } from "../../lib/formState.svelte.js";
    import { GetCorrespondents, AddCorrespondent } from "../../../wailsjs/go/src/WorkspaceBridge.js";

    let { correspondents = $bindable([]) } = $props();
    let showSettings = $state(false);
    let newName = $state("");

    $effect(() => {
        loadCorrespondents();
    });

    async function loadCorrespondents() {
        try {
            const list = await GetCorrespondents();
            if (list) correspondents = list;
        } catch (e) {
            console.error("Fehler beim Laden der Korrespondenten:", e);
        }
    }

    async function addCorrespondent() {
        const name = newName.trim();
        if (!name) return;
        try {
            await AddCorrespondent(name);
            newName = "";
            await loadCorrespondents();
        } catch (e) {
            console.error("Fehler beim Hinzufügen:", e);
        }
    }
</script>

<div
    class="p-4 rounded-2xl bg-zinc-900 border border-zinc-700 shadow-2xl space-y-2"
>
    <div class="flex justify-between items-center">
        <div class="text-sm font-semibold text-blue-400">Korrespondent</div>
        <button
            onclick={() => showSettings = !showSettings}
            class="text-zinc-400 hover:text-blue-300 transition-colors"
            aria-label="Korrespondenten verwalten"
        >
            <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M17 3a2.828 2.828 0 1 1 4 4L7.5 20.5 2 22l1.5-5.5L17 3z"/>
            </svg>
        </button>
    </div>

    <select
        bind:value={form.correspondent}
        class="bg-zinc-800 border border-zinc-600 text-base font-semibold text-blue-300 rounded-lg p-2 text-xs w-full shadow-[inset_2px_2px_5px_rgba(0,0,0,0.5)]"
    >
        <option value="">Bitte wählen</option>
        {#each correspondents as corr}
            <option value={corr}>{corr}</option>
        {/each}
    </select>

    <div class="text-xs font-semibold text-blue-400">
        Wird für den Filenamen verwendet
    </div>

    {#if showSettings}
        <div class="p-3 bg-zinc-800 rounded-lg border border-zinc-600 space-y-2">
            <div class="text-xs font-semibold text-zinc-400">Korrespondenten verwalten</div>
            {#each correspondents as corr}
                <div class="text-xs text-blue-300 font-mono">{corr}</div>
            {/each}
            <div class="flex gap-2 pt-1">
                <input
                    type="text"
                    bind:value={newName}
                    onkeydown={(e) => {
                        if (e.key === "Enter") {
                            e.preventDefault();
                            addCorrespondent();
                        }
                    }}
                    placeholder="Neuer Korrespondent"
                    class="flex-grow bg-zinc-900 border border-zinc-600 text-blue-300 rounded-lg p-2 text-xs shadow-[inset_2px_2px_5px_rgba(0,0,0,0.5)]"
                />
                <button
                    onclick={addCorrespondent}
                    class="bg-green-700 text-white px-3 py-1 rounded-lg text-xs font-semibold hover:bg-green-600"
                >
                    +
                </button>
            </div>
        </div>
    {/if}
</div>
