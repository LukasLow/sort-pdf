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
    <div class="text-lg font-bold text-primary mb-4">Ordner (+Archiv)</div>

    <div class="p-4 rounded-xl bg-muted border border-border-dim space-y-3">
        <div class="text-sm text-muted-foreground">
            Ordner im <span class="font-mono text-foreground">+Archiv</span>-Verzeichnis:
        </div>

        {#if folders.length === 0}
            <div class="text-sm text-muted-foreground italic">Keine Ordner vorhanden</div>
        {:else}
            <div class="space-y-1">
                {#each folders as f}
                    <div class="text-sm font-mono text-primary px-2 py-1 bg-card rounded-md">
                        {f}
                    </div>
                {/each}
            </div>
        {/if}

        <div class="flex gap-2 pt-2 border-t border-border-dim">
            <input
                type="text"
                bind:value={newName}
                onkeydown={(e) => { if (e.key === "Enter") { e.preventDefault(); create(); } }}
                placeholder="Neuer Ordnername"
                class="flex-grow bg-card border border-border-dim text-primary rounded-lg p-2 text-sm shadow-[inset_2px_2px_5px_rgba(0,0,0,0.5)]"
            />
            <button
                onclick={create}
                class="bg-primary text-background px-4 py-2 rounded-lg text-sm font-semibold hover:bg-primary-dark"
            >
                Erstellen
            </button>
        </div>
    </div>
</div>
