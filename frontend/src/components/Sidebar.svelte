<script>
    import ShowSidebarHeader from "./sidebar/ShowSidebarHeader.svelte";
    import ShowSidebarViews from "./sidebar/ShowSidebarViews.svelte";
    import ShowSidebarFooter from "./sidebar/ShowSidebarFooter.svelte";

    let { workDir, currentPdf, onSelect } = $props();
</script>

<div
    class="w-1/3 h-full p-5 flex flex-col gap-4 bg-black overflow-y-auto border-l border-zinc-900"
>
    <ShowSidebarHeader />

    <ShowSidebarViews {currentPdf} />

    <ShowSidebarFooter {workDir} {onSelect} />
</div>
