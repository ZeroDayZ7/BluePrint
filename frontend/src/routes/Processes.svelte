<script lang="ts">
  import { GetProcesses, KillProcess } from "../../wailsjs/go/main/App";
  import { deviceState } from "../lib/deviceState.svelte";
  import { untrack } from "svelte";
  import ListContainer from "../components/ListContainer.svelte";
  import Loader from "../components/Loader.svelte";
  import ActionButton from "../components/ActionButton.svelte";
  import IndexBadge from "../components/IndexBadge.svelte";

  let isLoading = $state(false);
  let searchQuery = $state("");

  // Pobieranie z cache
  let processes = $derived(
    deviceState.activeDevice
      ? deviceState.processesCache[deviceState.activeDevice.id] || []
      : [],
  );

  // Filtrowanie i sortowanie
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

    // Optymistyczne usunięcie z UI
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
  <div
    class="grid grid-cols-12 px-3 py-2 text-[10px] font-bold text-slate-500 uppercase tracking-tighter border-b border-slate-800/50 bg-slate-900/20 rounded-t-lg"
  >
    <div class="col-span-1">#</div>
    <div class="col-span-2">PID</div>
    <div class="col-span-4">Process Name</div>
    <div class="col-span-2 text-center">CPU</div>
    <div class="col-span-3 text-right">Actions</div>
  </div>

  <div
    class="h-[300px] overflow-y-auto custom-scrollbar border border-t-0 border-slate-800/50 rounded-b-lg bg-slate-900/10"
  >
    {#if isLoading && sortedProcesses.length === 0}
      <div class="flex h-full items-center justify-center">
        <Loader message="Syncing..." size="md" />
      </div>
    {:else}
      <div class="divide-y divide-slate-800/20">
        {#each sortedProcesses as proc, i}
          <div
            class="grid grid-cols-12 items-center px-3 py-1.5 hover:bg-slate-800/30 transition-all group"
          >
            <IndexBadge value={i + 1} class="col-span-1" />

            <IndexBadge value={proc.pid} class="col-span-2" />

            <div
              class="col-span-4 text-[11px] text-slate-300 truncate font-medium"
            >
              {proc.name}
            </div>

            <IndexBadge value="{proc.cpu}%" class="col-span-2 text-center" />

            <div class="col-span-3 flex justify-end gap-1">
              <button
                class="opacity-0 group-hover:opacity-100 p-1 text-slate-500 hover:text-blue-400 transition-all cursor-help"
                title="Process Info"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="12"
                  height="12"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2.5"
                >
                  <circle cx="12" cy="12" r="10" />
                  <path d="M12 16v-4" />
                  <path d="M12 8h.01" />
                </svg>
              </button>

              <ActionButton
                label="Kill"
                icon="zap"
                onclick={() => handleKill(proc.pid)}
              />
            </div>
          </div>
        {/each}
      </div>
    {/if}
  </div>
</ListContainer>
