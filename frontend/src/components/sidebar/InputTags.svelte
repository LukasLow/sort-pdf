<script>
    let tags = [];
    let newTag = "";

    function addTag() {
        if (newTag.trim() !== "") {
            tags = [...tags, newTag.trim()];
            newTag = "";
        }
    }

    function removeTag(tagToRemove) {
        tags = tags.filter((tag) => tag !== tagToRemove);
    }
</script>

<div
    class="p-4 rounded-2xl bg-zinc-900 border border-zinc-700 rounded-2xl shadow-2xl space-y-2"
>
    <div class="text-sm font-semibold text-blue-400">Tags</div>

    <div class="flex flex-wrap gap-2">
        {#each tags as tag}
            <span
                class="bg-zinc-700 text-base font-semibold text-blue-300 text-xs px-2 py-1 rounded-md flex items-center gap-1"
            >
                {tag}
                <button
                    on:click={() => removeTag(tag)}
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
            on:keydown={(e) => e.key === "Enter" && addTag()}
            placeholder="Neues Tag hinzufügen"
            class="flex-grow bg-zinc-800 border border-zinc-600 text-base font-semibold text-blue-300 rounded-lg p-2 text-xs shadow-[inset_2px_2px_5px_rgba(0,0,0,0.5)]"
        />
        <button
            on:click={addTag}
            class="bg-zinc-700 text-base font-semibold text-blue-300 px-3 py-1 rounded-lg text-xs hover:bg-zinc-600"
        >
            Hinzufügen
        </button>
    </div>

    <div class="text-xs font-semibold text-blue-400">
        Tags für die Organisation
    </div>
</div>
