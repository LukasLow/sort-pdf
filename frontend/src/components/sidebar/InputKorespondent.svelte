<script>
    import { form } from "../../lib/formState.svelte.js";
    import { GetCorrespondents, GetCorrespondentFolders } from "../../../wailsjs/go/src/WorkspaceBridge.js";

    let { correspondents = $bindable([]) } = $props();
    let folderMap = $state({});

    $effect(() => {
        loadAll();
    });

    async function loadAll() {
        try {
            const list = await GetCorrespondents();
            const map = await GetCorrespondentFolders();
            if (list) correspondents = list;
            if (map) folderMap = map;
        } catch (e) {
            console.error("Fehler beim Laden:", e);
        }
    }
</script>

<div
    class="p-4 rounded-2xl bg-zinc-900 border border-zinc-700 shadow-2xl space-y-2"
>
    <div class="text-sm font-semibold text-blue-400">Korrespondent</div>

    <select
        bind:value={form.correspondent}
        class="bg-zinc-800 border border-zinc-600 text-base font-semibold text-blue-300 rounded-lg p-2 text-xs w-full shadow-[inset_2px_2px_5px_rgba(0,0,0,0.5)]"
    >
        <option value="">Bitte wählen</option>
        {#each correspondents.filter(c => folderMap[c]) as corr}
            <option value={corr}>{corr}</option>
        {/each}
    </select>

    <div class="text-xs font-semibold text-blue-400">
        Wird für den Filenamen verwendet
    </div>
</div>
