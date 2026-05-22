<script>
  let { workDir, onSelect } = $props();
</script>

<div class="p-4 rounded-2xl bg-white/5 border border-white/10 backdrop-blur-xl shadow-lg space-y-2">
  <div class="text-[9px] text-blue-400 font-bold uppercase">
    Aktiver Ordner
  </div>

  <div class="text-[11px] font-mono text-zinc-400 break-all">
    {workDir}
  </div>

  <button
    onclick={onSelect}
    class="w-full bg-zinc-800 hover:bg-zinc-700 rounded-lg py-2 text-xs text-zinc-200"
  >
    Ordner ändern
  </button>

  <div class="text-[10px] text-zinc-600 text-center">
    sort-pdf v0.1.0
  </div>
</div>
