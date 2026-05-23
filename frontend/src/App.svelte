<script>
    import { onMount } from "svelte";
    import {
        CheckInitialWorkDir,
        SelectWorkingDirectory,
        GetNextPDF,
        EnsureWorkDirStructure,
    } from "../wailsjs/go/src/WorkspaceBridge";

    import WorkspaceSelector from "./components/WorkspaceSelector.svelte";
    import Sidebar from "./components/Sidebar.svelte";
    import PdfViewer from "./components/PdfViewer.svelte";

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
            try {
                const msg = await EnsureWorkDirStructure();
                if (msg) alert(msg);
            } catch (e) {
                console.error("Fehler beim Prüfen der Ordnerstruktur:", e);
            }
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
        <!-- Linke Seite: Sidebar -->
        <Sidebar {workDir} {currentPdf} onSelect={handleFolderSelection} onAction={loadNextPDF} />

        <!-- Rechte Seite: PDFJS-Viewer -->
        <PdfViewer filename={currentPdf} />
    </main>
{/if}
