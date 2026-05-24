<script>
    import SettingsWorkDirView from "./SettingsWorkDirView.svelte";
    import SettingsKorespondentenView from "./SettingsKorespondentenView.svelte";
    import SettingsOrdnerView from "./SettingsOrdnerView.svelte";
    import { GetVersion } from "../../../wailsjs/go/src/WorkspaceBridge.js";

    let { show, onClose, workDir, onSelect } = $props();

    let currentTab = $state("workdir");
    let version = $state("");

    const tabs = [
        { id: "workdir", label: "Arbeitsverzeichnis" },
        { id: "korespondenten", label: "Korrespondenten" },
        { id: "ordner", label: "Ordner" },
    ];

    $effect(() => {
        if (show && !version) {
            GetVersion().then((v) => version = v);
        }
    });
</script>

{#if show}
    <div
        class="fixed inset-0 z-50 flex items-center justify-center"
    >
        <div
            class="absolute inset-0 bg-black/60"
            role="button"
            tabindex="-1"
            onclick={onClose}
            onkeydown={(e) => { if (e.key === "Enter") onClose?.(); }}
        ></div>

        <div
            class="relative w-[800px] max-h-[80vh] bg-zinc-900 border border-zinc-700 rounded-2xl shadow-2xl flex overflow-hidden"
        >
            <div class="w-48 shrink-0 bg-zinc-950 p-4 flex flex-col gap-1 border-r border-zinc-800 h-full">
                <div class="text-xs font-semibold text-zinc-500 uppercase tracking-wider mb-2">
                    Einstellungen
                </div>
                {#each tabs as tab}
                    <button
                        onclick={() => currentTab = tab.id}
                        class="text-left px-3 py-2 rounded-lg text-sm font-semibold transition-colors {currentTab === tab.id ? 'bg-zinc-800 text-blue-300' : 'text-zinc-400 hover:text-zinc-200 hover:bg-zinc-800/50'}"
                    >
                        {tab.label}
                    </button>
                {/each}
                <div class="mt-auto pt-4 text-xs text-zinc-600">
                    {version}
                </div>
            </div>

            <div class="flex-1 p-6 overflow-y-auto">
                {#if currentTab === "workdir"}
                    <SettingsWorkDirView {workDir} {onSelect} />
                {:else if currentTab === "korespondenten"}
                    <SettingsKorespondentenView />
                {:else if currentTab === "ordner"}
                    <SettingsOrdnerView />
                {/if}
            </div>
        </div>

        <button
            onclick={onClose}
            class="absolute top-4 right-4 text-zinc-400 hover:text-white transition-colors"
            aria-label="Einstellungen schließen"
        >
            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M18 6 6 18M6 6l12 12"/>
            </svg>
        </button>
    </div>
{/if}
