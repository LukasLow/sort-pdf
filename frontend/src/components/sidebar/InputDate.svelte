<script>
  const years = Array.from({ length: 126 }, (_, i) => 2026 - i);
  const months = Array.from({ length: 12 }, (_, i) => i + 1);
  const days = Array.from({ length: 31 }, (_, i) => i + 1);

  let year = 2026;
  let month = 1;
  let day = 1;
</script>

<div class="p-4 rounded-2xl bg-white/5 border border-white/10 backdrop-blur-xl shadow-lg space-y-2">
  <div class="text-xs text-zinc-500">Datum (für Filename)</div>

  <div class="flex gap-2">
    <select class="bg-zinc-900 border border-zinc-700 rounded-lg p-2 text-xs w-full" bind:value={year}>
      {#each years as y}
        <option value={y}>{y}</option>
      {/each}
    </select>

    <select class="bg-zinc-900 border border-zinc-700 rounded-lg p-2 text-xs w-full" bind:value={month}>
      {#each months as m}
        <option value={m}>{m.toString().padStart(2,'0')}</option>
      {/each}
    </select>

    <select class="bg-zinc-900 border border-zinc-700 rounded-lg p-2 text-xs w-full" bind:value={day}>
      {#each days as d}
        <option value={d}>{d.toString().padStart(2,'0')}</option>
      {/each}
    </select>
  </div>

  <div class="text-[10px] text-zinc-600">
    Format: YYYY-MM-DD
  </div>
</div>
