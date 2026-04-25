<script lang="ts">
  import { deviceState } from "../lib/deviceState.svelte";
</script>

<div
  class="p-4 rounded-2xl transition-all duration-300 border
  {deviceState.isConnected
    ? 'bg-blue-600/10 border-blue-500/20 shadow-sm'
    : 'bg-slate-800/40 border-slate-800/60'}"
>
  <div
    class="text-[10px] text-slate-500 uppercase font-black tracking-[0.15em] mb-3"
  >
    Device Status
  </div>

  <div class="flex items-center gap-3">
    <div
      class="w-2.5 h-2.5 rounded-full transition-all duration-500
      {deviceState.isConnected
        ? 'bg-blue-400 shadow-[0_0_10px_rgba(96,165,250,0.5)]'
        : 'bg-slate-600'}"
    ></div>

    <span
      class="text-xs font-bold tracking-wide uppercase transition-colors
      {deviceState.isConnected ? 'text-blue-400' : 'text-slate-400'}"
    >
      {deviceState.displayStatus}
    </span>
  </div>
</div>
