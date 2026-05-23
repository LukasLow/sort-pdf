<script>
    import ShowPreview from "./ShowPreview.svelte";
    import ShowVersion from "./ShowVersion.svelte";

    let { onSettings } = $props();
</script>

<div class="mt-auto flex flex-col gap-2">
    <ShowPreview />
    <ShowVersion {onSettings} />
</div>
