<script lang="ts">
  import Button from "./components/Button.svelte";
  import Dashboard from "./routes/Dashboard.svelte";
  import Settings from "./routes/Settings.svelte";

  let activeTab = $state("dashboard");
  let deviceStatus = $state("Disconnected");
  let deviceName = $state("None");
  let logs = $state(["System initialized...", "Vite + Svelte 5 running..."]);

  async function refreshDevices() {
    logs = [...logs, "Scanning for ADB devices..."];
  }
</script>

<main
  class="flex h-screen bg-[#0b0e14] text-slate-200 overflow-hidden select-none"
>
  <nav class="w-64 bg-[#0d1117] border-r border-slate-800/60 flex flex-col p-5">
    <div class="flex items-center gap-3 mb-10 px-1">
      <div
        class="w-9 h-9 bg-blue-600 rounded-xl flex items-center justify-center font-black text-white shadow-lg shadow-blue-900/20"
      >
        A
      </div>
      <h1 class="text-lg font-bold tracking-tighter text-white uppercase">
        Commander
      </h1>
    </div>

    <div class="space-y-1.5 flex-1">
      {#each ["dashboard", "files", "settings"] as tab}
        <button
          onclick={() => (activeTab = tab)}
          class="w-full flex items-center gap-3 px-4 py-3 rounded-xl transition-all capitalize font-medium
          {activeTab === tab
            ? 'bg-blue-600/10 text-blue-400 border border-blue-500/20 shadow-sm'
            : 'text-slate-500 hover:text-slate-300 hover:bg-slate-800/40'}"
        >
          {tab}
        </button>
      {/each}
    </div>

    <div class="p-4 bg-slate-900/40 rounded-2xl border border-slate-800/60">
      <div
        class="text-[10px] text-slate-500 uppercase font-black tracking-widest mb-2"
      >
        Device Status
      </div>
      <div class="flex items-center gap-2.5">
        <div
          class="w-2.5 h-2.5 rounded-full {deviceStatus === 'Connected'
            ? 'bg-green-500 shadow-[0_0_8px_rgba(34,197,94,0.4)]'
            : 'bg-red-500'}"
        ></div>
        <span class="text-xs font-bold text-slate-300 tracking-wide uppercase"
          >{deviceStatus}</span
        >
      </div>
    </div>
  </nav>

  <section class="flex-1 flex flex-col">
    <header
      class="h-16 border-b border-slate-800/60 flex items-center justify-between px-8 bg-[#0b0e14]/80 backdrop-blur-xl"
    >
      <div class="text-xs font-medium text-slate-500 uppercase tracking-widest">
        Device: <span class="text-blue-400 ml-1 font-mono">{deviceName}</span>
      </div>
      <Button onclick={refreshDevices} class="text-xs px-5 py-2">
        Refresh ADB
      </Button>
    </header>

    <div class="flex-1 overflow-y-auto p-8 custom-scrollbar">
      {#if activeTab === "dashboard"}
        <Dashboard bind:logs />
      {:else if activeTab === "settings"}
        <Settings />
      {/if}
    </div>
  </section>
</main>
