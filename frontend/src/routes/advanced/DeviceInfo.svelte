<script lang="ts">
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
    Layers,
  } from "lucide-svelte";

  let isLoading = $state(false);
  let info = $state({
    model: "-",
    manufacturer: "-",
    androidVer: "-",
    apiLevel: "-",
    cpu: "-",
    battery: "-",
    serial: "-",
  });

  async function loadData(forceRefresh = false) {
    const device = deviceState.activeDevice;
    if (!device?.id) return;

    // Sprawdź cache (jeśli nie wymuszamy odświeżenia)
    if (!forceRefresh && deviceState.deviceInfoCache[device.id]) {
      info = deviceState.deviceInfoCache[device.id];
      return;
    }

    isLoading = true;
    try {
      const data = await GetDeviceInfo(device.id);
      if (data) {
        const newInfo = {
          model: data.model || "-",
          manufacturer: data.manufacturer || "-",
          androidVer: data.androidVer || "-",
          apiLevel: data.apiLevel || "-",
          cpu: data.cpu || "-",
          battery: data.battery || "-",
          serial: data.serial || "-",
        };

        info = newInfo;
        // Zapisz do globalnego stanu (cache)
        deviceState.deviceInfoCache[device.id] = newInfo;
      }
    } catch (err) {
      console.error("Błąd pobierania Device Info:", err);
    } finally {
      isLoading = false;
    }
  }

  // Reaguj na zmianę urządzenia
  $effect(() => {
    if (deviceState.isConnected && deviceState.activeDevice?.id) {
      loadData();
    }
  });
</script>

<div class="h-full overflow-y-auto custom-scrollbar p-4">
  {#if isLoading && info.model === "-"}
    <div class="flex h-full items-center justify-center">
      <Loader message="Reading system info..." size="lg" />
    </div>
  {:else}
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div class="flex flex-col gap-3">
        <div class="flex items-center justify-between ml-1">
          <h3
            class="text-[10px] font-bold text-slate-600 uppercase tracking-[0.2em]"
          >
            Hardware
          </h3>
          {#if isLoading}
            <span
              class="text-[9px] text-blue-500 animate-pulse font-bold uppercase"
              >Updating...</span
            >
          {/if}
        </div>

        <InfoCard
          label="Manufacturer"
          value={info.manufacturer}
          icon={Factory}
        />
        <InfoCard label="Model Name" value={info.model} icon={Cpu} />
        <InfoCard
          label="Processor / Architecture"
          value={info.cpu}
          icon={Layers}
        />
        <InfoCard label="Serial Number" value={info.serial} icon={Hash} />
      </div>

      <div class="flex flex-col gap-3">
        <h3
          class="text-[10px] font-bold text-slate-600 uppercase tracking-[0.2em] ml-1"
        >
          System & Power
        </h3>

        <InfoCard
          label="Battery Status"
          value={info.battery}
          variant={parseInt(info.battery) < 20 ? "amber" : "blue"}
          icon={BatteryCharging}
        />

        <InfoCard
          label="Android Version"
          value={`v${info.androidVer}`}
          subValue={`API LEVEL ${info.apiLevel}`}
          icon={Smartphone}
          variant="green"
        />

        <div
          class="mt-auto p-4 rounded-xl border border-dashed border-slate-800/50 flex flex-col items-center justify-center gap-2"
        >
          <span class="text-[9px] text-slate-500 uppercase font-medium"
            >Data is cached for performance</span
          >
          <button
            onclick={() => loadData(true)}
            class="text-[10px] px-3 py-1 bg-slate-800 hover:bg-slate-700 text-slate-300 rounded-lg transition-colors font-bold uppercase"
          >
            Force Refresh
          </button>
        </div>
      </div>
    </div>
  {/if}
</div>
