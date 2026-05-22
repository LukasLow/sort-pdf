<script lang="ts">
  import { onMount } from 'svelte';
  import { CheckInitialWorkDir, SelectWorkingDirectory } from '../wailsjs/go/src/WorkspaceBridge';

  import WorkspaceSelector from './components/WorkspaceSelector.svelte';
  import PdfViewer from './components/PdfViewer.svelte';
  import Sidebar from './components/Sidebar.svelte';

  let workDir = $state('');
  let currentPdf = $state('test.pdf');

  onMount(async () => {
    workDir = await CheckInitialWorkDir();
  });

  async function handleFolderSelection() {
    const selected = await SelectWorkingDirectory();
    if (selected) workDir = selected;
  }
</script>

{#if !workDir}
  <WorkspaceSelector onSelect={handleFolderSelection} />
{:else}
  <main class="flex h-screen w-screen bg-zinc-950 text-zinc-100 overflow-hidden antialiased font-sans">

    <!-- Linke Seite: 2/3 PDF Viewer -->
    <PdfViewer filename={currentPdf} />

    <!-- Rechte Seite: 1/3 Steuerung mit Pfad-Daten -->
    <Sidebar workDir={workDir} onSelect={handleFolderSelection} />

  </main>
{/if}
