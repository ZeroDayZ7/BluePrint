<script lang="ts">
  import { GetInstalledApps } from "../../wailsjs/go/main/App";
  import { deviceState } from "../lib/deviceState.svelte";

  let apps = $state<string[]>([]);
  let showUserApps = $state(true);

  async function loadApps() {
    if (!deviceState.activeDevice) return;
    apps = await GetInstalledApps(deviceState.activeDevice.id, showUserApps);
  }

  // Automatyczne przeładowanie przy zmianie opcji
  $effect(() => {
    if (deviceState.isConnected) loadApps();
  });
</script>

<div
  class="flex flex-col gap-4 p-4 bg-[#161b22]/40 rounded-2xl border border-slate-800/60"
>
  <div class="flex items-center justify-between">
    <h3 class="text-xs font-black uppercase tracking-widest text-slate-500">
      Applications
    </h3>

    <div class="flex bg-slate-900 rounded-lg p-1 border border-slate-800">
      <button
        class="px-3 py-1 text-[10px] rounded-md transition-all {showUserApps
          ? 'bg-blue-600 text-white'
          : 'text-slate-400'}"
        onclick={() => (showUserApps = true)}
      >
        USER
      </button>
      <button
        class="px-3 py-1 text-[10px] rounded-md transition-all {!showUserApps
          ? 'bg-blue-600 text-white'
          : 'text-slate-400'}"
        onclick={() => (showUserApps = false)}
      >
        SYSTEM
      </button>
    </div>
  </div>

  <div class="max-h-[400px] overflow-y-auto space-y-1 pr-2 custom-scrollbar">
    {#each apps as app}
      <div
        class="group flex items-center justify-between p-2 rounded-lg hover:bg-slate-800/50 transition-colors border border-transparent hover:border-slate-700/50"
      >
        <span class="text-xs text-slate-400 truncate">{app}</span>
        <button
          class="opacity-0 group-hover:opacity-100 text-[10px] text-red-400 font-bold uppercase tracking-tighter transition-opacity"
        >
          Uninstall
        </button>
      </div>
    {/each}
  </div>
</div>
