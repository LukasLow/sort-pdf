<script>
    import { form, getFilename } from "../lib/formState.svelte.js";
    import { MoveToArchiv, MoveToTodo, MoveToTrash, GetCorrespondentFolders, AnalyzePDF } from "../../wailsjs/go/src/WorkspaceBridge.js";

    import ShowSidebarHeader from "./sidebar/ShowSidebarHeader.svelte";
    import ShowSidebarViews from "./sidebar/ShowSidebarViews.svelte";
    import ShowSidebarFooter from "./sidebar/ShowSidebarFooter.svelte";
    import Settings from "./sidebar/Settings.svelte";

    let { workDir, currentPdf, onSelect, onAction } = $props();

    let showSettings = $state(false);

    $effect(() => {
        if (currentPdf) {
            analyzeCurrentPdf(currentPdf);
        }
    });

    async function analyzeCurrentPdf(fileName) {
        form.info = '';
        form.extras = '';
        try {
            const result = await AnalyzePDF(fileName);
            if (result) {
                console.log('Analyse-Ergebnis:', result);
                form.correspondent = result.correspondent || '';
            }
        } catch (e) {
            console.error("Fehler bei PDF-Analyse:", e);
        }
    }

    async function handleArchive() {
        if (!currentPdf) return;
        try {
            const folders = await GetCorrespondentFolders();
            const folder = folders[form.correspondent];
            if (!folder) {
                alert("Bitte in den Einstellungen einen Ordner für diesen Korrespondenten zuweisen.");
                return;
            }
            const targetName = getFilename();
            await MoveToArchiv(currentPdf, folder, targetName);
            onAction?.();
        } catch (e) {
            console.error("Fehler beim Archivieren:", e);
        }
    }

    async function handleTodo() {
        if (!currentPdf) return;
        try {
            const targetName = getFilename();
            await MoveToTodo(currentPdf, targetName);
            onAction?.();
        } catch (e) {
            console.error("Fehler:", e);
        }
    }

    async function handleTrash() {
        if (!currentPdf) return;
        try {
            const targetName = getFilename();
            await MoveToTrash(currentPdf, targetName);
            onAction?.();
        } catch (e) {
            console.error("Fehler:", e);
        }
    }

</script>

<div
    class="w-[275px] min-w-[275px] flex-shrink-0 h-full p-5 flex flex-col gap-4 bg-black overflow-y-auto border-r border-zinc-900"
>
    <ShowSidebarHeader />

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
