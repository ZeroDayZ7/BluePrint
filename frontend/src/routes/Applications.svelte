<script lang="ts">
  import { GetInstalledApps } from "../../wailsjs/go/main/App";
  import { deviceState } from "../lib/deviceState.svelte";
  import { untrack } from "svelte";
  import ListContainer from "../components/ListContainer.svelte";
  import Loader from "../components/Loader.svelte";
  import Checkbox from "../components/Checkbox.svelte";
  import IndexBadge from "../components/IndexBadge.svelte";
  import Button from "../components/Button.svelte";
  import { Trash2 } from "lucide-svelte";

  let showUserApps = $state(true);
  let isLoading = $state(false);
  let searchQuery = $state("");
  let debouncedQuery = $state("");
  let selectedPackages = $state<Set<string>>(new Set());

  $effect(() => {
    if (!searchQuery) {
      debouncedQuery = "";
      return;
    }
    const timeout = setTimeout(() => {
      debouncedQuery = searchQuery;
    }, 250);

    return () => clearTimeout(timeout);
  });

  // Filtrowanie listy
  let filteredApps = $derived.by(() => {
    const id = deviceState.activeDevice?.id;
    if (!id) return [];
    const baseList = showUserApps
      ? deviceState.userAppsCache[id] || []
      : deviceState.systemAppsCache[id] || [];

    if (!debouncedQuery) return baseList;

    const term = debouncedQuery.toLowerCase();
    return baseList.filter((pkg) => pkg.toLowerCase().includes(term));
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
    <div
      class="flex bg-slate-900 rounded-lg p-0.5 font-bold border border-slate-800"
    >
      <button
        class="px-2 py-0.5 text-xs rounded-md transition-all {showUserApps
          ? 'bg-blue-600 text-white'
          : 'text-slate-500 hover:text-slate-300'}"
        onclick={() => {
          showUserApps = true;
          selectedPackages.clear();
          loadApps();
        }}>USER</button
      >
      <button
        class="px-2 py-0.5 text-xs rounded-md transition-all {!showUserApps
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
      <Button
        onclick={handleBulkDelete}
        variant="danger"
        size="icon"
        title="Uninstall selected"
      >
        <Trash2 size={16} strokeWidth={2} />
      </Button>
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

            <Button
              variant="action"
              size="sm"
              onclick={() => handleSingleUninstall(app)}
            >
              <Trash2 size={12} strokeWidth={2.5} />
              <span>Uninstall</span>
            </Button>
          </div>
        {/each}
      </div>
    {/if}
  </div>
</ListContainer>
