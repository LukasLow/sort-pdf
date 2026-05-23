<script>
    import ShowPreview from "./ShowPreview.svelte";
    import ShowWorkDir from "./ShowWorkdir.svelte";
    import ShowVersion from "./ShowVersion.svelte";
    import { version } from "../../../package.json";

    let { workDir, onSelect } = $props();
</script>

<div class="mt-auto flex flex-col gap-2">
    <ShowPreview />
    <ShowWorkDir {workDir} {onSelect} />
    <ShowVersion {version} />
</div>
