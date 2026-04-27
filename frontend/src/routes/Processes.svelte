<script lang="ts">
  import { GetProcesses, KillProcess } from "../../wailsjs/go/main/App";
  import { deviceState } from "../lib/deviceState.svelte";
  import { untrack } from "svelte";
  import ListContainer from "../components/ListContainer.svelte";
  import Loader from "../components/Loader.svelte";
  import IndexBadge from "../components/IndexBadge.svelte";
  import Button from "../components/Button.svelte";
  import { Info, Zap } from "lucide-svelte";
  import Tabs from "../components/Tabs.svelte";
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
      console.error("Failed to fetch processes:", err);
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

  const processTabs = [
    { id: "user", label: "User" },
    { id: "system", label: "System" },
  ];

  let activeProcessType = $state("user");
  function handleTabChange(id: string) {
    activeProcessType = id;
  }
</script>

<ListContainer
  title="Running Processes"
  subtitle="{sortedProcesses.length} threads"
  searchPlaceholder="Filter by name or PID..."
  bind:searchQuery
  onRefresh={refreshProcesses}
>
  {#snippet headerExtra()}
    <div class="w-40">
      <Tabs
        tabs={processTabs}
        activeTab={activeProcessType}
        onChange={handleTabChange}
      />
    </div>
  {/snippet}

  <DataTable items={sortedProcesses}>
    {#snippet header()}
      <div
        class="grid grid-cols-12 px-3 py-2 text-xs font-bold text-slate-500 uppercase tracking-blueprint"
      >
        <div class="col-span-1">#</div>
        <div class="col-span-2">PID</div>
        <div class="col-span-4">Process Name</div>
        <div class="col-span-2 text-center">CPU</div>
        <div class="col-span-3 text-right pr-2">Actions</div>
      </div>
    {/snippet}

    {#snippet row(proc, i)}
      <div class="grid grid-cols-12 items-center px-3 transition-all">
        <IndexBadge value={i + 1} class="col-span-1" />
        <IndexBadge value={proc.pid} class="col-span-2" />
        <div class="col-span-4 text-xs truncate font-medium font-mono">
          {proc.name}
        </div>
        <IndexBadge value="{proc.cpu}%" class="col-span-2 text-center" />
        <div class="col-span-3 flex justify-end gap-1">
          <Button variant="action" size="icon" title="Info"
            ><Info size={12} /></Button
          >
          <Button
            variant="action"
            size="sm"
            onclick={() => handleKill(proc.pid)}
          >
            <Zap size={12} class="text-amber-400" fill="currentColor" />
            <span class="text-xs">Kill</span>
          </Button>
        </div>
      </div>
    {/snippet}
  </DataTable>
</ListContainer>
