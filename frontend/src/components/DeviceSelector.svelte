<script lang="ts">
  import { onMount } from "svelte";
  import { GetDevices } from "../../wailsjs/go/main/App";
  import { deviceState } from "../lib/deviceState.svelte";
  import Button from "./Button.svelte"; // Załóżmy taką nazwę pliku dla Twojego przycisku

  async function handleRefresh() {
    if (deviceState.isRefreshing) return;
    deviceState.isRefreshing = true;

    const minTime = new Promise((resolve) => setTimeout(resolve, 800));

    try {
      const result = await GetDevices();
      deviceState.devices = result || [];
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
    if (deviceState.devices.length === 0) handleRefresh();
  });
</script>

<div class="flex items-center gap-4 p-1.5 pr-3">
  <div class="flex flex-col px-3 min-w-[200px]">
    <span
      class="text-[9px] text-slate-500 uppercase tracking-[0.2em] font-black leading-none mb-1.5"
    >
      Active Device
    </span>

    <div class="flex items-center gap-2 h-5">
      <div
        class="w-2 h-2 rounded-full transition-all duration-500 {deviceState.isConnected
          ? 'bg-blue-400 shadow-[0_0_10px_rgba(96,165,250,0.8)]'
          : 'bg-slate-700'}"
      ></div>

      <div class="flex-1 truncate">
        {#if deviceState.devices.length > 1}
          <select
            bind:value={deviceState.selectedDeviceIndex}
            class="bg-transparent text-[13px] font-mono border-none text-blue-400 focus:outline-none cursor-pointer w-full"
          >
            {#each deviceState.devices as device, i}
              <option value={i} class="bg-slate-900 text-white">
                {device.model}
              </option>
            {/each}
          </select>
        {:else if deviceState.devices.length === 1}
          <span class="text-[13px] font-mono text-blue-400 truncate block">
            {deviceState.activeDevice?.model}
          </span>
        {:else}
          <span class="text-[13px] font-mono text-slate-500 italic"
            >No Device Found</span
          >
        {/if}
      </div>
    </div>
  </div>

  <Button
    onclick={handleRefresh}
    variant="secondary"
    class="!p-2.5 !rounded-xl active:scale-95"
    disabled={deviceState.isRefreshing}
    title="Refresh ADB devices"
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="2.5"
      stroke="currentColor"
      class="w-4 h-4 {deviceState.isRefreshing
        ? 'animate-spin text-blue-400'
        : 'text-slate-400'}"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99"
      />
    </svg>
  </Button>
</div>
