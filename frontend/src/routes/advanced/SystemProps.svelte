<script lang="ts">
  import ListContainer from "../../components/ListContainer.svelte";
  import Loader from "../../components/Loader.svelte";
  import { GetDeviceInfo } from "../../../wailsjs/go/main/App";
  import { deviceState } from "../../lib/deviceState.svelte";
  import DataTable from "../../components/DataTable.svelte";

  let searchQuery = $state("");
  let isLoading = $state(false);
  let allProps = $state([] as { key: string; value: string }[]);

  let filteredProps = $derived(
    allProps.filter(
      (p) =>
        p.key.toLowerCase().includes(searchQuery.toLowerCase()) ||
        p.value.toLowerCase().includes(searchQuery.toLowerCase()),
    ),
  );

  async function loadData(forceRefresh = false) {
    const device = deviceState.activeDevice;
    if (!device?.id) return;

    if (!forceRefresh && deviceState.systemPropsCache[device.id]) {
      allProps = deviceState.systemPropsCache[device.id];
      return;
    }

    isLoading = true;
    try {
      const data = await GetDeviceInfo(device.id);
      if (data && data.allProps) {
        allProps = data.allProps;
        deviceState.systemPropsCache[device.id] = data.allProps;
      }
    } catch (err) {
      console.error(err);
    } finally {
      isLoading = false;
    }
  }

  $effect(() => {
    if (deviceState.isConnected && deviceState.activeDevice?.id) {
      loadData();
    }
  });
</script>

<div class="flex flex-col h-full relative">
  {#if isLoading && allProps.length === 0}
    <div class="flex-1 flex items-center justify-center">
      <Loader message="Fetching props..." size="lg" />
    </div>
  {:else}
    <ListContainer
      title="System Properties"
      subtitle="Raw getprop output ({filteredProps.length} entries)"
      searchPlaceholder="Filter props..."
      bind:searchQuery
      onRefresh={() => loadData(true)}
    >
      <DataTable items={filteredProps} height="h-[calc(100vh-280px)]">
        {#snippet header()}
          <div
            class="grid grid-cols-2 px-4 py-2 text-xs font-bold text-slate-500 uppercase tracking-blueprint"
          >
            <div>Key</div>
            <div>Value</div>
          </div>
        {/snippet}

        {#snippet row(prop)}
          <div class="grid grid-cols-2 text-xs font-mono">
            <div
              class="px-4 py-1.5 text-slate-500 group-hover:text-blue-400 break-all border-r border-slate-800/10"
            >
              {prop.key}
            </div>
            <div class="px-4 py-1.5 text-slate-300 break-all">
              {prop.value}
            </div>
          </div>
        {/snippet}
      </DataTable>
    </ListContainer>
  {/if}
</div>
