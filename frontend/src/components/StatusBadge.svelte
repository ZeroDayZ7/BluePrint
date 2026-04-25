<script lang="ts">
  let { status = "Disconnected" } = $props();
  let isConnected = $derived(status === "Connected");
</script>

<div class="p-4 bg-slate-900/40 rounded-2xl border border-slate-800/60">
  <div
    class="text-[10px] text-slate-500 uppercase font-black tracking-widest mb-2"
  >
    Device Status
  </div>
  <div class="flex items-center gap-2.5">
    <div
      class="w-2.5 h-2.5 rounded-full {isConnected
        ? 'bg-green-500 shadow-[0_0_8px_rgba(34,197,94,0.4)]'
        : 'bg-red-500'}"
    ></div>
    <span class="text-xs font-bold text-slate-300 tracking-wide uppercase"
      >{status}</span
    >
  </div>
</div>
