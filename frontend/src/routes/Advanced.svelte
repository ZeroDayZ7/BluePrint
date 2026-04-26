<script lang="ts">
  import LogcatView from "./advanced/LogcatView.svelte";
  import PowerMenu from "./advanced/PowerMenu.svelte";
  import TerminalView from "./advanced/TerminalView.svelte";
  import DeviceInfo from "./advanced/DeviceInfo.svelte"; // Import nowej zakładki

  let activeTab = $state("terminal");

  const tabs = [
    { id: "terminal", label: "Terminal", icon: "terminal" },
    { id: "logcat", label: "Logcat", icon: "list" },
    { id: "device", label: "Device Info", icon: "info" }, // Nowa zakładka
    { id: "power", label: "Power Menu", icon: "power" },
  ];
</script>

<div class="flex flex-col gap-4 h-full">
  <div
    class="flex items-center gap-1 bg-slate-900/50 p-1 rounded-xl border border-slate-800/50 w-fit"
  >
    {#each tabs as tab}
      <button
        onclick={() => (activeTab = tab.id)}
        class="flex items-center gap-2 px-4 py-2 rounded-lg text-sm font-medium transition-all
        {activeTab === tab.id
          ? 'bg-blue-600 text-white shadow-lg'
          : 'text-slate-400 hover:text-slate-200 hover:bg-slate-800'}"
      >
        {tab.label}
      </button>
    {/each}
  </div>

  <div
    class="flex-1 bg-slate-900/20 border border-slate-800/40 rounded-2xl p-4 overflow-hidden relative"
  >
    {#if activeTab === "terminal"}
      <TerminalView />
    {:else if activeTab === "logcat"}
      <LogcatView />
    {:else if activeTab === "device"}
      <DeviceInfo />
    {:else if activeTab === "power"}
      <PowerMenu />
    {/if}
  </div>
</div>
