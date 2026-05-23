<script>
    import { form } from "../lib/formState.svelte.js";
    import { MoveToArchiv, MoveToTodo, MoveToTrash, GetCorrespondentFolders } from "../../wailsjs/go/src/WorkspaceBridge.js";

    import ShowSidebarHeader from "./sidebar/ShowSidebarHeader.svelte";
    import ShowSidebarViews from "./sidebar/ShowSidebarViews.svelte";
    import ShowSidebarFooter from "./sidebar/ShowSidebarFooter.svelte";
    import Settings from "./sidebar/Settings.svelte";

    let { workDir, currentPdf, onSelect, onAction } = $props();

    let showSettings = $state(false);

    async function handleArchive() {
        if (!currentPdf) return;
        try {
            const folders = await GetCorrespondentFolders();
            const folder = folders[form.correspondent];
            if (!folder) {
                alert("Bitte in den Einstellungen einen Ordner für diesen Korrespondenten zuweisen.");
                return;
            }
            await MoveToArchiv(currentPdf, folder);
            onAction?.();
        } catch (e) {
            console.error("Fehler beim Archivieren:", e);
        }
    }

    async function handleTodo() {
        if (!currentPdf) return;
        try {
            await MoveToTodo(currentPdf);
            onAction?.();
        } catch (e) {
            console.error("Fehler:", e);
        }
    }

    async function handleTrash() {
        if (!currentPdf) return;
        try {
            await MoveToTrash(currentPdf);
            onAction?.();
        } catch (e) {
            console.error("Fehler:", e);
        }
    }
</script>

<div
    class="w-1/3 h-full p-5 flex flex-col gap-4 bg-black overflow-y-auto border-l border-zinc-900"
>
    <ShowSidebarHeader onSettings={() => showSettings = true} />

    <ShowSidebarViews
        {currentPdf}
        onArchive={handleArchive}
        onTodo={handleTodo}
        onTrash={handleTrash}
    />

    <ShowSidebarFooter onSettings={() => showSettings = true} />
</div>

<Settings
    show={showSettings}
    onClose={() => showSettings = false}
    {workDir}
    {onSelect}
/>
