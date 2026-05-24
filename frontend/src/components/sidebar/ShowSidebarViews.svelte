<script>
    import ShowPdfInfo from "./ShowPdfInfo.svelte";
    import InputDate from "./InputDate.svelte";
    import InputKorespondent from "./InputKorespondent.svelte";
    import InputInfo from "./InputInfo.svelte";
    import InputExtras from "./InputExtras.svelte";
    import ActionButton from "./ActionButton.svelte";

    let { currentPdf, onArchive, onTodo, onTrash } = $props();

    let currentView = $state(0);
    const totalViews = 4;

    function canProceed() {
        return true;
    }
</script>

<div class="flex flex-col gap-4">
    {#if currentView === 0}
        <ShowPdfInfo {currentPdf} />
    {:else if currentView === 1}
        <InputDate />
    {:else if currentView === 2}
        <InputKorespondent />
        <InputInfo />
        <InputExtras />
    {:else if currentView === 3}
        <ActionButton {onArchive} {onTodo} {onTrash} />
    {/if}

    <div class="flex justify-between items-center gap-4">
        <button
            onclick={() => currentView = Math.max(0, currentView - 1)}
            disabled={currentView === 0}
            class="flex-1 px-4 py-2 bg-zinc-800 rounded-lg text-sm font-semibold text-zinc-300 hover:bg-zinc-700 disabled:opacity-50 disabled:cursor-not-allowed"
        >
            Zurück
        </button>
        <span class="text-xs text-zinc-500 shrink-0">
            {currentView + 1} / {totalViews}
        </span>
        <button
            onclick={() => currentView = Math.min(totalViews - 1, currentView + 1)}
            disabled={currentView === totalViews - 1 || !canProceed()}
            class="flex-1 px-4 py-2 bg-zinc-800 rounded-lg text-sm font-semibold text-zinc-300 hover:bg-zinc-700 disabled:opacity-50 disabled:cursor-not-allowed"
        >
            Weiter
        </button>
    </div>
</div>
