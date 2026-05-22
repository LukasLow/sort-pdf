<script>
  let { correspondents } = $props();
</script>

<div class="p-4 rounded-2xl bg-white/5 border border-white/10 backdrop-blur-xl shadow-lg space-y-2">
  <div class="text-xs text-zinc-500">Korrespondent</div>

  <select class="w-full bg-zinc-900 border border-zinc-700 rounded-lg p-2 text-xs text-zinc-200">
    <option value="">-- wählen --</option>

    {#each correspondents as c}
      <option value={c}>{c}</option>
    {/each}
  </select>

  <div class="text-[10px] text-zinc-600">
    → wird Teil vom Filename
  </div>
</div>
