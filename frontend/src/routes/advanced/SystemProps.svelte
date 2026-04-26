<script lang="ts">
  import ListContainer from "../../components/ListContainer.svelte";
  import Loader from "../../components/Loader.svelte";
  import { GetDeviceInfo } from "../../../wailsjs/go/main/App";
  import { deviceState } from "../../lib/deviceState.svelte";
  import PropertyTable from "../../components/PropertyTable.svelte";

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

    // Sprawdź cache
    if (!forceRefresh && deviceState.systemPropsCache[device.id]) {
      allProps = deviceState.systemPropsCache[device.id];
      return;
    }

    isLoading = true;
    try {
      const data = await GetDeviceInfo(device.id);
      if (data && data.allProps) {
        allProps = data.allProps;
        deviceState.systemPropsCache[device.id] = data.allProps; // Zapisz do cache
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

<div class="flex flex-col h-full overflow-hidden relative">
  {#if isLoading && allProps.length === 0}
    <div class="flex-1 flex items-center justify-center">
      <Loader message="Fetching props..." size="lg" />
    </div>
  {:else}
    <div class="flex-1 overflow-y-auto custom-scrollbar">
      <ListContainer
        title="System Properties"
        subtitle="Raw getprop output ({filteredProps.length} entries)"
        searchPlaceholder="Filter props..."
        bind:searchQuery
        onRefresh={() => loadData(true)}
      >
        <PropertyTable items={filteredProps} />
      </ListContainer>
    </div>
  {/if}
</div>
