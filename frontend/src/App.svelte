<script>
    import { onMount } from "svelte";
    import {
        CheckInitialWorkDir,
        SelectWorkingDirectory,
        GetNextPDF,
    } from "../wailsjs/go/src/WorkspaceBridge";

    import WorkspaceSelector from "./components/WorkspaceSelector.svelte";
    import PdfViewer from "./components/PdfViewer.svelte";
    import Sidebar from "./components/Sidebar.svelte";

    let workDir = $state("");
    let currentPdf = $state(""); // Startet jetzt komplett leer

    // Holt die älteste PDF aus dem Go-Backend
    async function loadNextPDF() {
        try {
            currentPdf = await GetNextPDF();
        } catch (err) {
            console.error("Fehler beim Laden der nächsten PDF:", err);
        }
    }

    onMount(async () => {
        workDir = await CheckInitialWorkDir();
        if (workDir) {
            await loadNextPDF();
        }
    });

    async function handleFolderSelection() {
        const selected = await SelectWorkingDirectory();
        if (selected) {
            workDir = selected;
            await loadNextPDF();
        }
    }
</script>

{#if !workDir}
    <WorkspaceSelector onSelect={handleFolderSelection} />
{:else}
    <main
        class="flex h-screen w-screen bg-zinc-950 text-zinc-100 overflow-hidden antialiased font-sans"
    >
        <!-- Linke Seite: PDFJS-Viewer zeigt die dynamisch ermittelte Datei -->
        <PdfViewer filename={currentPdf} />

        <!-- Rechte Seite: Sidebar kriegt die Ordner-Wechsel-Funktion -->
        <Sidebar {workDir} {currentPdf} onSelect={handleFolderSelection} />
    </main>
{/if}
