<script>
    import { GetFolderList, CreateFolder } from "../../../wailsjs/go/src/WorkspaceBridge.js";

    let folders = $state([]);
    let newName = $state("");

    $effect(() => {
        load();
    });

    async function load() {
        try {
            folders = await GetFolderList();
        } catch (e) {
            console.error("Fehler beim Laden der Ordner:", e);
        }
    }

    async function create() {
        const name = newName.trim();
        if (!name) return;
        try {
            await CreateFolder(name);
            newName = "";
            await load();
        } catch (e) {
            console.error("Fehler beim Erstellen:", e);
        }
    }
</script>

<div>
    <div class="text-lg font-bold text-blue-300 mb-4">Ordner (+Archiv)</div>

    <div class="p-4 rounded-xl bg-zinc-800 border border-zinc-700 space-y-3">
        <div class="text-sm text-zinc-400">
            Ordner im <span class="font-mono text-zinc-300">+Archiv</span>-Verzeichnis:
        </div>

        {#if folders.length === 0}
            <div class="text-sm text-zinc-500 italic">Keine Ordner vorhanden</div>
        {:else}
            <div class="space-y-1">
                {#each folders as f}
                    <div class="text-sm font-mono text-blue-300 px-2 py-1 bg-zinc-900 rounded-md">
                        {f}
                    </div>
                {/each}
            </div>
        {/if}

        <div class="flex gap-2 pt-2 border-t border-zinc-700">
            <input
                type="text"
                bind:value={newName}
                onkeydown={(e) => { if (e.key === "Enter") { e.preventDefault(); create(); } }}
                placeholder="Neuer Ordnername"
                class="flex-grow bg-zinc-900 border border-zinc-600 text-blue-300 rounded-lg p-2 text-sm shadow-[inset_2px_2px_5px_rgba(0,0,0,0.5)]"
            />
            <button
                onclick={create}
                class="bg-blue-700 text-white px-4 py-2 rounded-lg text-sm font-semibold hover:bg-blue-600"
            >
                Erstellen
            </button>
        </div>
    </div>
</div>
