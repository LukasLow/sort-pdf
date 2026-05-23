<script>
    import ShowPreview from "./ShowPreview.svelte";
    import SettingsButton from "./SettingsButton.svelte";

    let { onSettings } = $props();
</script>

<div class="mt-auto flex flex-col gap-2">
    <ShowPreview />
    <SettingsButton {onSettings} />
</div>
