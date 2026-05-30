<script>
    import { form } from "../../lib/formState.svelte.js";
    import { GetCorrespondents, GetCorrespondentFolders } from "../../../wailsjs/go/src/WorkspaceBridge.js";

    let correspondents = $state([]);
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
    class="p-4 rounded-2xl bg-card border border-border shadow-2xl space-y-2"
>
    <div class="text-sm font-semibold text-primary">Korrespondent</div>

    <select
        bind:value={form.correspondent}
        class="bg-muted border border-border-dim text-base font-semibold text-primary rounded-lg p-2 text-xs w-full shadow-[inset_2px_2px_5px_rgba(0,0,0,0.5)]"
    >
        <option value="">Bitte wählen</option>
        {#each correspondents as corr}
            <option value={corr}>{corr}{folderMap[corr] ? "" : " (kein Ordner)"}</option>
        {/each}
    </select>

    <div class="text-xs font-semibold text-primary">
        Wird für den Filenamen verwendet
    </div>
</div>
