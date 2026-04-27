<script lang="ts">
  let { logs = [] } = $props();
  let isExpanded = $state(false);

  function toggleConsole() {
    isExpanded = !isExpanded;
  }
</script>

<div
  class="border-t border-slate-800/60 bg-[#0b0e14]/80 backdrop-blur-xl transition-all duration-300 ease-in-out overflow-hidden flex flex-col shrink-0"
  style="height: {isExpanded ? '300px' : '40px'}"
>
  <button
    type="button"
    class="w-full flex items-center justify-between px-6 py-2 shrink-0 hover:bg-slate-800/30 transition-colors focus:outline-none focus:bg-slate-800/50"
    onclick={toggleConsole}
    aria-expanded={isExpanded}
    aria-controls="console-logs"
  >
    <div class="flex items-center gap-3">
      <div
        class="text-[10px] font-black uppercase tracking-[0.2em] text-slate-500"
      >
        Output Log
      </div>
      <div
        class="px-1.5 py-0.5 rounded bg-blue-500/10 text-blue-400 text-xs font-bold"
      >
        {logs.length} lines
      </div>
    </div>

    <div class="text-slate-500 transition-colors group-hover:text-blue-400">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="w-4 h-4 transition-transform duration-300 {isExpanded
          ? 'rotate-180'
          : ''}"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
        aria-hidden="true"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M5 15l7-7 7 7"
        />
      </svg>
    </div>
  </button>

  <div
    id="console-logs"
    class="flex-1 overflow-y-auto px-6 py-2 font-mono text-[11px] space-y-1 custom-scrollbar select-text"
  >
    {#each logs as log}
      <div class="flex gap-3 border-b border-slate-900/50 pb-1">
        <span class="text-slate-600 shrink-0"
          >[{new Date().toLocaleTimeString()}]</span
        >
        <span class="text-blue-500/80 shrink-0 tracking-tighter font-bold"
          >system:</span
        >
        <span class="text-slate-300 break-all">{log}</span>
      </div>
    {/each}
  </div>
</div>
