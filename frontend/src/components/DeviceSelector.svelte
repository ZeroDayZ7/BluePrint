<script lang="ts">
  import { onMount } from "svelte";
  import { GetDevices } from "../../wailsjs/go/main/App";
  import { deviceState } from "../lib/deviceState.svelte";

  async function handleRefresh() {
    if (deviceState.isRefreshing) return;
    deviceState.isRefreshing = true;

    // Timer dla płynności animacji kręcenia się ikony
    const minTime = new Promise((resolve) => setTimeout(resolve, 1000));

    try {
      const result = await GetDevices();
      deviceState.devices = result || [];

      // Jeśli po odświeżeniu lista jest mniejsza niż wybrany indeks, zresetuj go
      if (deviceState.selectedDeviceIndex >= deviceState.devices.length) {
        deviceState.selectedDeviceIndex = 0;
      }
    } catch (err) {
      console.error("ADB Error:", err);
    } finally {
      await minTime;
      deviceState.isRefreshing = false;
    }
  }

  onMount(() => {
    if (deviceState.devices.length === 0) {
      handleRefresh();
    }
  });
</script>

<div class="flex items-center gap-6">
  <div class="flex flex-col items-start">
    <span
      class="text-[10px] text-slate-500 uppercase tracking-[0.2em] font-bold leading-none mb-2"
    >
      Active Device
    </span>

    <div class="flex items-center gap-2">
      <div
        class="w-2 h-2 rounded-full transition-colors {deviceState.isConnected
          ? 'bg-blue-400 shadow-[0_0_8px_rgba(96,165,250,0.5)]'
          : 'bg-slate-600'}"
      ></div>

      {#if deviceState.devices.length > 1}
        <select
          bind:value={deviceState.selectedDeviceIndex}
          class="bg-transparent text-sm font-mono border-none text-blue-400 focus:ring-0 cursor-pointer outline-none p-0"
        >
          {#each deviceState.devices as device, i}
            <option value={i} class="bg-slate-900 text-white">
              {device.model} ({device.id})
            </option>
          {/each}
        </select>
      {:else if deviceState.devices.length === 1}
        <span
          class="text-sm font-mono {deviceState.isConnected
            ? 'text-blue-400'
            : 'text-slate-400'}"
        >
          {deviceState.activeDevice?.model} ({deviceState.activeDevice?.id})
        </span>
      {:else}
        <span class="text-sm font-mono text-slate-400"> No Device Found </span>
      {/if}
    </div>
  </div>

  <button
    onclick={handleRefresh}
    disabled={deviceState.isRefreshing}
    class="p-2 rounded-lg bg-slate-800/50 border border-slate-700/50 hover:bg-slate-700 hover:border-slate-600 transition-all group active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed"
    title="Refresh ADB devices"
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="2"
      stroke="currentColor"
      class="w-4 h-4 {deviceState.isRefreshing ? 'animate-spin' : ''}"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99"
      />
    </svg>
  </button>
</div>
