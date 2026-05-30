<script>
    import { form, getFilename } from "../lib/formState.svelte.js";
    import {
        MoveToArchiv, MoveToArchivOverwrite,
        MoveToTodo, MoveToTodoOverwrite,
        MoveToTrash, MoveToTrashOverwrite,
        CheckArchiveConflict, CheckTodoConflict, CheckTrashConflict,
        GetCorrespondentFolders, AnalyzePDF
    } from "../../wailsjs/go/src/WorkspaceBridge.js";

    import ShowSidebarHeader from "./sidebar/ShowSidebarHeader.svelte";
    import ShowSidebarViews from "./sidebar/ShowSidebarViews.svelte";
    import ShowSidebarFooter from "./sidebar/ShowSidebarFooter.svelte";
    import Settings from "./sidebar/Settings.svelte";
    import MoveConflictDialog from "./sidebar/MoveConflictDialog.svelte";

    let { workDir, currentPdf, onSelect, onAction } = $props();

    let showSettings = $state(false);

    // Zustand für Dateikonflikt-Dialog
    let showConflictDialog = $state(false);
    let conflictData = $state(null);
    let conflictAction = $state(null); // 'archive' | 'todo' | 'trash'
    let conflictSubFolder = $state('');

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

    function openMoveConflictDialog(action, conflict, subFolder) {
        conflictAction = action;
        conflictData = conflict;
        conflictSubFolder = subFolder || '';
        showConflictDialog = true;
    }

    async function handleOverwrite() {
        if (!currentPdf || !conflictData) return;
        showConflictDialog = false;
        try {
            const targetName = getFilename();
            if (conflictAction === 'archive') {
                await MoveToArchivOverwrite(currentPdf, conflictSubFolder, targetName);
            } else if (conflictAction === 'todo') {
                await MoveToTodoOverwrite(currentPdf, targetName);
            } else if (conflictAction === 'trash') {
                await MoveToTrashOverwrite(currentPdf, targetName);
            }
            onAction?.();
        } catch (e) {
            console.error("Fehler beim Überschreiben:", e);
            alert("Fehler beim Überschreiben: " + (e.message || e));
        }
    }

    function handleCancelConflict() {
        showConflictDialog = false;
        conflictData = null;
        conflictAction = null;
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
            const conflict = await CheckArchiveConflict(currentPdf, folder, targetName);
            if (conflict.hasConflict) {
                openMoveConflictDialog('archive', conflict, folder);
                return;
            }
            await MoveToArchiv(currentPdf, folder, targetName);
            onAction?.();
        } catch (e) {
            console.error("Fehler beim Archivieren:", e);
            alert("Fehler beim Archivieren: " + (e.message || e));
        }
    }

    async function handleTodo() {
        if (!currentPdf) return;
        try {
            const targetName = getFilename();
            const conflict = await CheckTodoConflict(currentPdf, targetName);
            if (conflict.hasConflict) {
                openMoveConflictDialog('todo', conflict);
                return;
            }
            await MoveToTodo(currentPdf, targetName);
            onAction?.();
        } catch (e) {
            console.error("Fehler:", e);
            alert("Fehler beim Verschieben zu Todo: " + (e.message || e));
        }
    }

    async function handleTrash() {
        if (!currentPdf) return;
        try {
            const targetName = getFilename();
            const conflict = await CheckTrashConflict(currentPdf, targetName);
            if (conflict.hasConflict) {
                openMoveConflictDialog('trash', conflict);
                return;
            }
            await MoveToTrash(currentPdf, targetName);
            onAction?.();
        } catch (e) {
            console.error("Fehler:", e);
            alert("Fehler beim Verschieben in den Papierkorb: " + (e.message || e));
        }
    }

</script>

<div
    class="w-[275px] min-w-[275px] flex-shrink-0 h-full p-5 flex flex-col gap-4 bg-background overflow-y-auto border-r border-border-dim"
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

<MoveConflictDialog
    show={showConflictDialog}
    conflict={conflictData}
    onOverwrite={handleOverwrite}
    onCancel={handleCancelConflict}
/>
