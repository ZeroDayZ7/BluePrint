<script lang="ts">
  import { GetInstalledApps } from "../../wailsjs/go/main/App";
  import { deviceState } from "../lib/deviceState.svelte";
  import { untrack } from "svelte";
  import Loader from "../components/Loader.svelte";
  import Checkbox from "../components/Checkbox.svelte";
  import SearchBar from "../components/SearchBar.svelte";
  import IconButton from "../components/IconButton.svelte";
  import ActionButton from "../components/ActionButton.svelte";

  let showUserApps = $state(true);
  let isLoading = $state(false);
  let searchQuery = $state("");
  let selectedPackages = $state<Set<string>>(new Set());

  // Filtrowanie listy
  let filteredApps = $derived.by(() => {
    const id = deviceState.activeDevice?.id;
    if (!id) return [];
    const baseList = showUserApps
      ? deviceState.userAppsCache[id] || []
      : deviceState.systemAppsCache[id] || [];

    if (!searchQuery) return baseList;
    return baseList.filter((pkg) =>
      pkg.toLowerCase().includes(searchQuery.toLowerCase()),
    );
  });

  async function loadApps(force = false) {
    // 1. Sprawdź czy już nie ładujemy (bardzo ważne!)
    if (isLoading) return;

    const device = deviceState.activeDevice;
    if (!device) return;

    // 2. Sprawdź cache
    const currentCache = showUserApps
      ? deviceState.userAppsCache[device.id]
      : deviceState.systemAppsCache[device.id];

    if (currentCache?.length > 0 && !force) return;

    isLoading = true;
    try {
      const result = await GetInstalledApps(device.id, showUserApps);
      // Używamy untrack przy zapisie do cache, żeby nie triggerować efektu
      untrack(() => {
        if (showUserApps) deviceState.userAppsCache[device.id] = result;
        else deviceState.systemAppsCache[device.id] = result;
      });
    } catch (err) {
      console.error("ADB Error:", err);
    } finally {
      isLoading = false;
    }
  }

  function toggleSelect(pkg: string) {
    if (selectedPackages.has(pkg)) selectedPackages.delete(pkg);
    else selectedPackages.add(pkg);
    selectedPackages = new Set(selectedPackages);
  }

  function handleSingleUninstall(pkg: string) {
    console.log("Uninstalling single:", pkg);
  }

  function handleBulkDelete() {
    if (selectedPackages.size === 0) return;
    console.log("Bulk uninstall:", Array.from(selectedPackages));
  }

  $effect(() => {
    const id = deviceState.activeDevice?.id;
    const connected = deviceState.isConnected;
    const type = showUserApps;

    if (connected && id) {
      untrack(() => loadApps());
    }
  });
</script>

<div
  class="flex flex-col gap-4 p-4 bg-[#161b22]/40 rounded-2xl border border-slate-800/60"
>
  <div class="flex items-center justify-between">
    <div class="flex flex-col">
      <h3 class="text-xs font-black uppercase tracking-widest text-slate-500">
        Applications
      </h3>
      <span class="text-[10px] text-slate-600 font-medium"
        >{filteredApps.length} packages</span
      >
    </div>

    <div class="flex bg-slate-900 rounded-lg p-1 border border-slate-800">
      <button
        class="px-3 py-1 text-[10px] rounded-md transition-all {showUserApps
          ? 'bg-blue-600 text-white'
          : 'text-slate-400'}"
        onclick={() => {
          showUserApps = true;
          selectedPackages.clear();
        }}>USER</button
      >
      <button
        class="px-3 py-1 text-[10px] rounded-md transition-all {!showUserApps
          ? 'bg-blue-600 text-white'
          : 'text-slate-400'}"
        onclick={() => {
          showUserApps = false;
          selectedPackages.clear();
        }}>SYSTEM</button
      >
    </div>
  </div>

  <div class="flex items-center gap-2">
    <SearchBar bind:value={searchQuery} placeholder="Filter apps..." />

    <div class="flex items-center gap-1">
      <div
        class={selectedPackages.size === 0
          ? "opacity-30 grayscale pointer-events-none"
          : ""}
      >
        <IconButton
          onclick={handleBulkDelete}
          variant="danger"
          title="Uninstall selected"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
            ><path d="M3 6h18" /><path
              d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6"
            /><path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2" /><line
              x1="10"
              y1="11"
              x2="10"
              y2="17"
            /><line x1="14" y1="11" x2="14" y2="17" /></svg
          >
        </IconButton>
      </div>

      <IconButton onclick={() => loadApps(true)} title="Refresh list">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="16"
          height="16"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
          ><path d="M21 12a9 9 0 1 1-9-9c2.52 0 4.93 1 6.74 2.74L21 8" /><path
            d="M21 3v5h-5"
          /></svg
        >
      </IconButton>
    </div>
  </div>

  <div class="h-[300px] overflow-y-auto pr-2 custom-scrollbar">
    {#if isLoading}
      <div class="flex h-64 items-center justify-center">
        <Loader message="Scanning..." size="md" />
      </div>
    {:else}
      <div class="space-y-1">
        {#each filteredApps as app}
          <div
            class="group flex items-center gap-3 p-2 rounded-lg hover:bg-slate-800/40 border border-transparent hover:border-slate-800 transition-all"
          >
            <Checkbox
              checked={selectedPackages.has(app)}
              onchange={() => toggleSelect(app)}
            />

            <span
              class="text-[11px] text-slate-400 truncate text-left flex-1 font-mono"
            >
              {app}
            </span>

            <ActionButton
              label="Uninstall"
              icon="trash"
              onclick={() => handleSingleUninstall(app)}
            />
          </div>
        {/each}
      </div>
    {/if}
  </div>
</div>
