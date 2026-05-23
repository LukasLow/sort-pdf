<script>
    let { archive, todo, trash } = $props();
</script>

<div
    class="p-4 rounded-2xl bg-zinc-900 border border-zinc-700 rounded-2xl shadow-2xl flex justify-between items-center"
>
    <button
        on:click={archive}
        class="bg-blue-600 text-blue-300 px-4 py-2 rounded-lg text-sm font-semibold hover:bg-blue-700"
    >
        Archivieren
    </button>
    <button
        on:click={todo}
        class="bg-yellow-600 text-blue-300 px-4 py-2 rounded-lg text-sm font-semibold hover:bg-yellow-700"
    >
        ToDo
    </button>
    <button
        on:click={trash}
        class="bg-red-600 text-blue-300 px-4 py-2 rounded-lg text-sm font-semibold hover:bg-red-700"
    >
        Papierkorb
    </button>
</div>
