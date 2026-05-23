<script>
    let { onArchive, onTodo, onTrash } = $props();
</script>

<div
    class="p-4 rounded-2xl bg-zinc-900 border border-zinc-700 shadow-2xl flex flex-col gap-3"
>
    <button
        onclick={onArchive}
        class="w-full bg-blue-600 text-white px-4 py-3 rounded-lg text-sm font-semibold hover:bg-blue-700"
    >
        Archivieren
    </button>
    <button
        onclick={onTodo}
        class="w-full bg-yellow-600 text-white px-4 py-3 rounded-lg text-sm font-semibold hover:bg-yellow-700"
    >
        ToDo
    </button>
    <button
        onclick={onTrash}
        class="w-full bg-red-600 text-white px-4 py-3 rounded-lg text-sm font-semibold hover:bg-red-700"
    >
        Papierkorb
    </button>
</div>
