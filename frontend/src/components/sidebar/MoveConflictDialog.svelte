<script>
    import { OpenFileInOSViewer } from "../../../wailsjs/go/src/WorkspaceBridge.js";

    let { show, conflict, onOverwrite, onCancel } = $props();

    async function openSource() {
        try {
            await OpenFileInOSViewer(conflict.sourcePath);
        } catch (e) {
            console.error("Fehler beim Öffnen der Quelldatei:", e);
        }
    }

    async function openTarget() {
        try {
            await OpenFileInOSViewer(conflict.targetPath);
        } catch (e) {
            console.error("Fehler beim Öffnen der Zieldatei:", e);
        }
    }
</script>

{#if show}
<div class="fixed inset-0 z-50 flex items-center justify-center">
    <div
        class="absolute inset-0 bg-black/60"
        role="button"
        tabindex="-1"
        onclick={onCancel}
        onkeydown={(e) => { if (e.key === "Enter") onCancel?.(); }}
    ></div>

    <div
        class="relative w-[480px] bg-card border border-border rounded-2xl shadow-2xl p-6 space-y-5"
    >
        <div class="text-lg font-bold text-danger">Datei existiert bereits</div>

        <div class="text-sm text-muted-foreground space-y-1">
            <div>Am Zielort existiert bereits eine Datei mit dem Namen:</div>
            <div class="font-mono text-primary font-semibold break-all bg-muted px-3 py-2 rounded-lg mt-2">
                {conflict.targetName}
            </div>
        </div>

        <div class="flex gap-3">
            <button
                onclick={openSource}
                class="flex-1 bg-muted text-foreground px-4 py-3 rounded-lg text-sm font-semibold hover:bg-primary-dark border border-border-dim"
            >
                Aktuelle PDF öffnen
            </button>
            <button
                onclick={openTarget}
                class="flex-1 bg-muted text-foreground px-4 py-3 rounded-lg text-sm font-semibold hover:bg-primary-dark border border-border-dim"
            >
                Vorhandene PDF öffnen
            </button>
        </div>

        <div class="flex gap-3 pt-2">
            <button
                onclick={onCancel}
                class="flex-1 bg-muted text-foreground px-4 py-3 rounded-lg text-sm font-semibold hover:bg-border-dim"
            >
                Abbrechen
            </button>
            <button
                onclick={onOverwrite}
                class="flex-1 bg-danger text-background px-4 py-3 rounded-lg text-sm font-semibold hover:opacity-85"
            >
                Überschreiben
            </button>
        </div>
    </div>
</div>
{/if}
