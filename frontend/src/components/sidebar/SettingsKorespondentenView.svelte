<script>
    import { GetCorrespondents, AddCorrespondent, RemoveCorrespondent, ReorderCorrespondents, GetCorrespondentFolders, SetCorrespondentFolder, GetFolderList } from "../../../wailsjs/go/src/WorkspaceBridge.js";

    let correspondents = $state([]);
    let folders = $state([]);
    let folderMap = $state({});
    let newName = $state("");

    $effect(() => {
        loadAll();
    });

    async function loadAll() {
        try {
            correspondents = await GetCorrespondents();
            folderMap = await GetCorrespondentFolders();
            folders = await GetFolderList();
        } catch (e) {
            console.error("Fehler beim Laden:", e);
        }
    }

    async function add() {
        const name = newName.trim();
        if (!name) return;
        try {
            await AddCorrespondent(name);
            newName = "";
            await loadAll();
        } catch (e) {
            console.error("Fehler:", e);
        }
    }

    async function remove(name) {
        try {
            await RemoveCorrespondent(name);
            await loadAll();
        } catch (e) {
            console.error("Fehler:", e);
        }
    }

    async function setFolder(correspondent, folder) {
        try {
            await SetCorrespondentFolder(correspondent, folder);
            folderMap = await GetCorrespondentFolders();
        } catch (e) {
            console.error("Fehler:", e);
        }
    }

    async function moveUp(index) {
        if (index === 0) return;
        const updated = [...correspondents];
        [updated[index - 1], updated[index]] = [updated[index], updated[index - 1]];
        try {
            await ReorderCorrespondents(updated);
            await loadAll();
        } catch (e) {
            console.error("Fehler:", e);
        }
    }

    async function moveDown(index) {
        if (index === correspondents.length - 1) return;
        const updated = [...correspondents];
        [updated[index], updated[index + 1]] = [updated[index + 1], updated[index]];
        try {
            await ReorderCorrespondents(updated);
            await loadAll();
        } catch (e) {
            console.error("Fehler:", e);
        }
    }
</script>

<div>
    <div class="text-lg font-bold text-blue-300 mb-4">Korrespondenten</div>

    <div class="space-y-3">
        <div class="flex gap-2">
            <input
                type="text"
                bind:value={newName}
                onkeydown={(e) => { if (e.key === "Enter") { e.preventDefault(); add(); } }}
                placeholder="Neuer Korrespondent"
                class="flex-grow bg-zinc-800 border border-zinc-600 text-blue-300 rounded-lg p-2 text-sm shadow-[inset_2px_2px_5px_rgba(0,0,0,0.5)]"
            />
            <button
                onclick={add}
                class="bg-green-700 text-white px-4 py-2 rounded-lg text-sm font-semibold hover:bg-green-600"
            >
                Hinzufügen
            </button>
        </div>

        {#each correspondents as corr, i}
            <div class="p-3 rounded-xl bg-zinc-800 border border-zinc-700 space-y-2">
                <div class="flex justify-between items-center">
                    <div class="flex items-center gap-2">
                        <div class="flex flex-col gap-0.5">
                            <button
                                onclick={() => moveUp(i)}
                                disabled={i === 0}
                                class="text-zinc-500 hover:text-blue-300 disabled:opacity-30 disabled:cursor-not-allowed leading-none"
                                aria-label="Nach oben"
                            >
                                <svg xmlns="http://www.w3.org/2000/svg" class="w-3 h-3" viewBox="0 0 24 24" fill="currentColor"><path d="M12 5l-7 7h14l-7-7z"/></svg>
                            </button>
                            <button
                                onclick={() => moveDown(i)}
                                disabled={i === correspondents.length - 1}
                                class="text-zinc-500 hover:text-blue-300 disabled:opacity-30 disabled:cursor-not-allowed leading-none"
                                aria-label="Nach unten"
                            >
                                <svg xmlns="http://www.w3.org/2000/svg" class="w-3 h-3" viewBox="0 0 24 24" fill="currentColor"><path d="M12 19l7-7H5l7 7z"/></svg>
                            </button>
                        </div>
                        <span class="text-sm font-semibold text-blue-300">{corr}</span>
                    </div>
                    <button
                        onclick={() => remove(corr)}
                        class="text-zinc-500 hover:text-red-400 transition-colors"
                        aria-label="Korrespondent entfernen"
                    >
                        <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <path d="M3 6h18M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
                        </svg>
                    </button>
                </div>
                <div class="flex items-center gap-2">
                    <span class="text-xs text-zinc-500 shrink-0">Ordner:</span>
                    <select
                        class="flex-grow bg-zinc-900 border border-zinc-600 text-blue-300 rounded-lg p-1.5 text-xs"
                        value={folderMap[corr] || ""}
                        onchange={(e) => setFolder(corr, e.target.value)}
                    >
                        <option value="">– kein Ordner –</option>
                        {#each folders as f}
                            <option value={f}>{f}</option>
                        {/each}
                    </select>
                </div>
            </div>
        {/each}
    </div>
</div>
