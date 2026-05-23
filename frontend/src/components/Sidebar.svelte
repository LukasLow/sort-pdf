<script>
    import { form } from "../lib/formState.svelte.js";
    import { MoveToArchiv, MoveToTodo, MoveToTrash, GetCorrespondentFolders, AnalyzePDF, WritePDFTags } from "../../wailsjs/go/src/WorkspaceBridge.js";

    import ShowSidebarHeader from "./sidebar/ShowSidebarHeader.svelte";
    import ShowSidebarViews from "./sidebar/ShowSidebarViews.svelte";
    import ShowSidebarFooter from "./sidebar/ShowSidebarFooter.svelte";
    import Settings from "./sidebar/Settings.svelte";

    let { workDir, currentPdf, onSelect, onAction } = $props();

    let showSettings = $state(false);
    let analysis = $state(null);

    $effect(() => {
        if (currentPdf) {
            analyzeCurrentPdf(currentPdf);
        }
    });

    async function analyzeCurrentPdf(fileName) {
        form.info = '';
        form.extras = '';
        form.tags = [];
        try {
            const result = await AnalyzePDF(fileName);
            analysis = result;
            if (result) {
                form.year = result.dateYear || new Date().getFullYear();
                form.month = result.dateMonth || new Date().getMonth() + 1;
                form.day = result.dateDay || new Date().getDate();
                form.correspondent = result.correspondent || '';
            }
        } catch (e) {
            console.error("Fehler bei PDF-Analyse:", e);
            analysis = null;
        }
    }

    async function writeTags() {
        try {
            await WritePDFTags(currentPdf, form.tags);
        } catch (e) {
            console.error("Fehler beim Schreiben der Tags:", e);
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
            await writeTags();
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
            await writeTags();
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
            await writeTags();
            const targetName = getFilename();
            await MoveToTrash(currentPdf, targetName);
            onAction?.();
        } catch (e) {
            console.error("Fehler:", e);
        }
    }

    function getFilename() {
        const pad = (n) => String(n).padStart(2, '0');
        const date = `${form.year}-${pad(form.month)}-${pad(form.day)}`;
        const parts = [date, form.correspondent, form.info, form.extras].filter(Boolean);
        return parts.join('_') + '.pdf';
    }
</script>

<div
    class="w-[275px] min-w-[275px] flex-shrink-0 h-full p-5 flex flex-col gap-4 bg-black overflow-y-auto border-r border-zinc-900"
>
    <ShowSidebarHeader onSettings={() => showSettings = true} />

    <ShowSidebarViews
        {currentPdf}
        {analysis}
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
