<script>
    let { onArchive, onTodo, onTrash } = $props();
</script>

<div
    class="p-4 rounded-2xl bg-card border border-border shadow-2xl flex flex-col gap-3"
>
    <button
        onclick={onArchive}
        class="w-full bg-primary text-background px-4 py-3 rounded-lg text-sm font-semibold hover:bg-primary-dark"
    >
        Archivieren
    </button>
    <button
        onclick={onTodo}
        class="w-full bg-accent text-background px-4 py-3 rounded-lg text-sm font-semibold hover:opacity-85"
    >
        ToDo
    </button>
    <button
        onclick={onTrash}
        class="w-full bg-danger text-background px-4 py-3 rounded-lg text-sm font-semibold hover:opacity-85"
    >
        Papierkorb
    </button>
</div>
