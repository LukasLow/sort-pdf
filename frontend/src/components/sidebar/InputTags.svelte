<script>
    import { form } from "../../lib/formState.svelte.js";

    let newTag = $state("");

    function addTag() {
        if (newTag.trim() !== "") {
            form.tags = [...form.tags, newTag.trim()];
            newTag = "";
        }
    }

    function removeTag(tagToRemove) {
        form.tags = form.tags.filter((tag) => tag !== tagToRemove);
    }
</script>

<div
    class="p-4 rounded-2xl bg-zinc-900 border border-zinc-700 shadow-2xl space-y-2"
>
    <div class="text-sm font-semibold text-blue-400">Tags</div>

    <div class="flex flex-wrap gap-2">
        {#each form.tags as tag}
            <span
                class="bg-zinc-700 text-base font-semibold text-blue-300 text-xs px-2 py-1 rounded-md flex items-center gap-1"
            >
                {tag}
                <button
                    onclick={() => removeTag(tag)}
                    class="text-blue-400 hover:text-white"
                >
                    &times;
                </button>
            </span>
        {/each}
    </div>

    <div class="flex gap-2">
        <input
            type="text"
            bind:value={newTag}
            onkeydown={(e) => {
                if (e.key === "Enter") {
                    e.preventDefault();
                    addTag();
                }
            }}
            placeholder="Neues Tag hinzufügen"
            class="flex-grow bg-zinc-800 border border-zinc-600 text-base font-semibold text-blue-300 rounded-lg p-2 text-xs shadow-[inset_2px_2px_5px_rgba(0,0,0,0.5)]"
        />
        <button
            onclick={addTag}
            class="bg-zinc-700 text-base font-semibold text-blue-300 px-3 py-1 rounded-lg text-xs hover:bg-zinc-600"
        >
            Hinzufügen
        </button>
    </div>
</div>
