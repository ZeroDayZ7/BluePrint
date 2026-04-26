<script lang="ts">
  import ListContainer from "../../components/ListContainer.svelte";
  import Loader from "../../components/Loader.svelte";
  import { GetDeviceInfo } from "../../../wailsjs/go/main/App";
  import { deviceState } from "../../lib/deviceState.svelte";
  import InfoCard from "../../components/InfoCard.svelte";
  import {
    Cpu,
    Smartphone,
    BatteryCharging,
    Factory,
    Hash,
  } from "lucide-svelte";

  let searchQuery = $state("");
  let isLoading = $state(false);

  let info = $state({
    model: "-",
    manufacturer: "-",
    androidVer: "-",
    apiLevel: "-",
    cpu: "-",
    battery: "-",
    serial: "-",
    allProps: [] as { key: string; value: string }[],
  });

  let filteredProps = $derived(
    info.allProps.filter(
      (p) =>
        p.key.toLowerCase().includes(searchQuery.toLowerCase()) ||
        p.value.toLowerCase().includes(searchQuery.toLowerCase()),
    ),
  );

  async function loadData() {
    // POBIERAMY ID Z GLOBALNEGO STANU
    const device = deviceState.activeDevice;

    if (!device?.id) {
      console.warn("JS: Brak aktywnego urządzenia w deviceState.");
      return;
    }

    console.log("JS: Próba wywołania GetDeviceInfo dla:", device.id);
    isLoading = true;
    try {
      const data = await GetDeviceInfo(device.id);
      if (data) {
        // Mapowanie pól (upewnij się, że nazwy pól z Go zgadzają się z JS - Wails zwykle zmienia pierwszą literę na małą)
        info = {
          model: data.model || "-",
          manufacturer: data.manufacturer || "-",
          androidVer: data.androidVer || "-",
          apiLevel: data.apiLevel || "-",
          cpu: data.cpu || "-",
          battery: data.battery || "-",
          serial: data.serial || "-",
          allProps: data.allProps || [],
        };
        console.log("JS: Dane załadowane pomyślnie");
      }
    } catch (err) {
      console.error("Błąd backendu Go:", err);
    } finally {
      isLoading = false;
    }
  }

  // Reaguj na zmianę aktywnego urządzenia
  $effect(() => {
    if (deviceState.isConnected && deviceState.activeDevice?.id) {
      loadData();
    }
  });
</script>

<div class="flex flex-col gap-4 h-full">
  {#if isLoading && info.allProps.length === 0}
    <div class="flex-1 flex items-center justify-center">
      <Loader message="Fetching Device Intelligence..." size="lg" />
    </div>
  {:else}
    <div class="grid grid-cols-2 lg:grid-cols-5 gap-3 shrink-0">
      <InfoCard
        label="Battery Status"
        value={info.battery}
        variant={parseInt(info.battery) < 20 ? "amber" : "blue"}
        icon={BatteryCharging}
      />

      <InfoCard label="Manufacturer" value={info.manufacturer} icon={Factory} />

      <InfoCard
        label="System Version"
        value={`v${info.androidVer}`}
        subValue={`SDK ${info.apiLevel}`}
        icon={Smartphone}
      />

      <InfoCard
        label="Model Name"
        value={info.model}
        variant="default"
        icon={Cpu}
      />

      <InfoCard label="Serial Number" value={info.serial} icon={Hash} />
    </div>

    <div class="flex-1 min-h-0 relative">
      {#if isLoading}
        <div
          class="absolute inset-0 z-50 bg-slate-950/40 backdrop-blur-sm rounded-2xl flex items-center justify-center"
        >
          <Loader message="Updating properties..." size="md" />
        </div>
      {/if}

      <ListContainer
        title="System Properties"
        subtitle="Raw getprop output ({filteredProps.length} entries)"
        searchPlaceholder="Filter by key or value..."
        bind:searchQuery
        onRefresh={loadData}
      >
        <div
          class="h-full overflow-y-auto custom-scrollbar rounded-xl border border-slate-800/50 bg-[#0b0e14]/50"
        >
          <table class="w-full text-left text-[11px] border-collapse relative">
            <thead
              class="sticky top-0 bg-[#161b22] border-b border-slate-800 z-10"
            >
              <tr>
                <th
                  class="px-4 py-2.5 font-bold text-slate-500 uppercase tracking-tighter w-1/2"
                  >Key</th
                >
                <th
                  class="px-4 py-2.5 font-bold text-slate-500 uppercase tracking-tighter w-1/2"
                  >Value</th
                >
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-800/50">
              {#each filteredProps as prop}
                <tr class="hover:bg-blue-500/5 transition-colors group">
                  <td
                    class="px-4 py-2 font-mono text-slate-500 group-hover:text-blue-400 break-all border-r border-slate-800/30"
                  >
                    {prop.key}
                  </td>
                  <td class="px-4 py-2 font-mono text-slate-300 break-all">
                    {prop.value}
                  </td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      </ListContainer>
    </div>
  {/if}
</div>
