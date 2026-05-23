<script>
    import ShowSidebarHeader from "./sidebar/ShowSidebarHeader.svelte";
    import ShowPdfInfo from "./sidebar/ShowPdfInfo.svelte";
    import InputDate from "./sidebar/InputDate.svelte";
    import InputKorespondent from "./sidebar/InputKorespondent.svelte";
    import InputInfo from "./sidebar/InputInfo.svelte";
    import InputExtras from "./sidebar/InputExtras.svelte";
    import InputTags from "./sidebar/InputTags.svelte";
    import ShowPreview from "./sidebar/ShowPreview.svelte";
    import ActionButton from "./sidebar/ActionButton.svelte";
    import ShowSidebarFooter from "./sidebar/ShowSidebarFooter.svelte";

    let { workDir, currentPdf, onSelect } = $props();

    function handleAction(type) {
        console.log("ACTION:", type);
    }
</script>

<div
    class="w-1/3 h-full p-5 flex flex-col gap-4 bg-black overflow-y-auto border-l border-zinc-900"
>
    <ShowSidebarHeader />

    <ShowPdfInfo {currentPdf} />

    <InputDate />

    <InputKorespondent correspondents={[]} />

    <InputInfo />

    <InputExtras />

    <InputTags />

    <ShowPreview />

    <ActionButton
        archive={() => handleAction("archive")}
        todo={() => handleAction("todo")}
        trash={() => handleAction("trash")}
    />

    <ShowSidebarFooter {workDir} {onSelect} />
</div>
