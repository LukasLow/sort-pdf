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
    let currentPdf = $state("");
    let emptyMessage = $state("");

    async function loadNextPDF() {
        try {
            const pdf = await GetNextPDF();
            if (pdf) {
                currentPdf = pdf;
                emptyMessage = "";
            } else {
                currentPdf = "";
                emptyMessage = "Keine PDFs im Eingang";
            }
        } catch (err) {
            console.error("Fehler beim Laden der nächsten PDF:", err);
            emptyMessage = "Fehler beim Laden der PDF";
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

        <!-- Rechte Seite -->
        {#if emptyMessage}
            <div class="flex-1 flex items-center justify-center bg-zinc-900 text-zinc-400 text-sm">
                {emptyMessage}
            </div>
        {:else}
            <PdfViewer filename={currentPdf} />
        {/if}
    </main>
{/if}
