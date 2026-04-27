<script lang="ts">
  import { GetProcesses, KillProcess } from "../../wailsjs/go/main/App";
  import { deviceState } from "../lib/deviceState.svelte";
  import { untrack } from "svelte";
  import ListContainer from "../components/ListContainer.svelte";
  import Loader from "../components/Loader.svelte";
  import IndexBadge from "../components/IndexBadge.svelte";
  import Button from "../components/Button.svelte";
  import { Info, Zap } from "lucide-svelte";
  import DataTable from "../components/DataTable.svelte";

  let isLoading = $state(false);
  let searchQuery = $state("");

  let processes = $derived(
    deviceState.activeDevice
      ? deviceState.processesCache[deviceState.activeDevice.id] || []
      : [],
  );

  let sortedProcesses = $derived(
    processes
      .filter(
        (p) =>
          p.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
          p.pid.includes(searchQuery),
      )
      .sort((a, b) => a.name.localeCompare(b.name)),
  );

  async function refreshProcesses() {
    if (!deviceState.activeDevice || isLoading) return;
    isLoading = true;
    try {
      const data = await GetProcesses(deviceState.activeDevice.id);
      untrack(() => {
        if (deviceState.activeDevice) {
          deviceState.processesCache[deviceState.activeDevice.id] = data || [];
        }
      });
    } catch (err) {
      console.error(err);
    } finally {
      isLoading = false;
    }
  }

  async function handleKill(pid: string) {
    if (!deviceState.activeDevice) return;
    const deviceId = deviceState.activeDevice.id;
    const previous = [...deviceState.processesCache[deviceId]];

    deviceState.processesCache[deviceId] = previous.filter(
      (p) => p.pid !== pid,
    );

    try {
      const result = await KillProcess(deviceId, pid);
      if (result !== "Success") {
        deviceState.processesCache[deviceId] = previous;
        alert("Error: " + result);
      }
    } catch (err) {
      deviceState.processesCache[deviceId] = previous;
    }
  }

  $effect(() => {
    if (deviceState.isConnected && deviceState.activeDevice) {
      if (!deviceState.processesCache[deviceState.activeDevice.id]) {
        untrack(() => refreshProcesses());
      }
    }
  });
</script>

<ListContainer
  title="Running Processes"
  subtitle="{sortedProcesses.length} threads"
  searchPlaceholder="Filter by name or PID..."
  bind:searchQuery
  onRefresh={refreshProcesses}
>
  {#if isLoading && sortedProcesses.length === 0}
    <div class="flex h-64 items-center justify-center">
      <Loader message="Fetching processes..." size="md" />
    </div>
  {:else}
    <DataTable items={sortedProcesses} height="h-[calc(100vh-300px)]">
      {#snippet header()}
        <div
          class="grid grid-cols-12 px-4 py-2.5 text-[10px] font-black text-slate-500 uppercase tracking-widest"
        >
          <div class="col-span-1">#</div>
          <div class="col-span-2">PID</div>
          <div class="col-span-5">Process Name</div>
          <div class="col-span-2 text-center">CPU</div>
          <div class="col-span-2 text-right">Actions</div>
        </div>
      {/snippet}

      {#snippet row(proc, i)}
        <div class="grid grid-cols-12 items-center px-4 py-1.5">
          <div class="col-span-1">
            <IndexBadge value={i + 1} />
          </div>
          <div class="col-span-2">
            <span class="text-[11px] font-mono text-blue-400/80"
              >{proc.pid}</span
            >
          </div>
          <div
            class="col-span-5 text-slate-300 text-xs truncate font-mono tracking-tight cursor-help"
            title={proc.name}
          >
            {proc.name}
          </div>
          <div class="col-span-2 text-center">
            <span
              class="text-[11px] font-mono {parseFloat(proc.cpu) > 50
                ? 'text-red-400'
                : 'text-emerald-400'}"
            >
              {proc.cpu}%
            </span>
          </div>
          <div
            class="col-span-2 flex justify-end gap-1 opacity-0 group-hover:opacity-100 transition-opacity"
          >
            <Button variant="action" size="icon" title="Info">
              <Info size={12} />
            </Button>
            <Button
              variant="action"
              size="sm"
              onclick={() => handleKill(proc.pid)}
            >
              <Zap size={12} class="text-amber-400" fill="currentColor" />
              <span class="text-[10px] font-bold uppercase">Kill</span>
            </Button>
          </div>
        </div>
      {/snippet}
    </DataTable>
  {/if}
</ListContainer>
