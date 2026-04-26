<script lang="ts">
  import { GetInstalledApps } from "../../wailsjs/go/main/App";
  import { deviceState } from "../lib/deviceState.svelte";
  import { untrack } from "svelte";
  import ListContainer from "../components/ListContainer.svelte";
  import Loader from "../components/Loader.svelte";
  import Checkbox from "../components/Checkbox.svelte";
  import IconButton from "../components/IconButton.svelte";
  import ActionButton from "../components/ActionButton.svelte";
  import IndexBadge from "../components/IndexBadge.svelte";

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
    if (isLoading) return;
    const device = deviceState.activeDevice;
    if (!device) return;

    const currentCache = showUserApps
      ? deviceState.userAppsCache[device.id]
      : deviceState.systemAppsCache[device.id];

    if (currentCache?.length > 0 && !force) return;

    isLoading = true;
    try {
      const result = await GetInstalledApps(device.id, showUserApps);
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
    if (deviceState.isConnected && id) {
      untrack(() => loadApps());
    }
  });
</script>

<ListContainer
  title="Applications"
  subtitle="{filteredApps.length} packages"
  searchPlaceholder="Filter apps..."
  bind:searchQuery
  onRefresh={() => loadApps(true)}
>
  {#snippet headerExtra()}
    <div class="flex bg-slate-900 rounded-lg p-0.5 border border-slate-800">
      <button
        class="px-2 py-0.5 text-[9px] rounded-md transition-all {showUserApps
          ? 'bg-blue-600 text-white'
          : 'text-slate-500 hover:text-slate-300'}"
        onclick={() => {
          showUserApps = true;
          selectedPackages.clear();
          loadApps();
        }}>USER</button
      >
      <button
        class="px-2 py-0.5 text-[9px] rounded-md transition-all {!showUserApps
          ? 'bg-blue-600 text-white'
          : 'text-slate-500 hover:text-slate-300'}"
        onclick={() => {
          showUserApps = false;
          selectedPackages.clear();
          loadApps();
        }}>SYSTEM</button
      >
    </div>
  {/snippet}

  {#snippet searchActions()}
    <div
      class={selectedPackages.size === 0
        ? "opacity-20 grayscale pointer-events-none"
        : "transition-all"}
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
          ><path d="M3 6h18" /><path
            d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6"
          /><path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2" /></svg
        >
      </IconButton>
    </div>
  {/snippet}

  <div class="h-[300px] overflow-y-auto pr-2 custom-scrollbar">
    {#if isLoading && filteredApps.length === 0}
      <div class="flex h-full items-center justify-center">
        <Loader message="Scanning device..." size="md" />
      </div>
    {:else}
      <div class="space-y-1">
        {#each filteredApps as app, i}
          <div
            class="group flex items-center gap-3 p-2 rounded-lg hover:bg-slate-800/40 border border-transparent hover:border-slate-800 transition-all"
          >
            <IndexBadge value={i + 1} class="w-4" />

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
</ListContainer>
