<script>
    import ShowPreview from "./ShowPreview.svelte";
    import ShowWorkDir from "./ShowWorkdir.svelte";
    import ShowVersion from "./ShowVersion.svelte";
    import { version } from "../../../package.json";

    let { previewUrl, workDir, onSelect } = $props();
</script>

<div class="mt-auto flex flex-col gap-2">
    <ShowPreview {previewUrl} />
    <ShowWorkDir {workDir} {onSelect} />
    <ShowVersion {version} />
</div>
