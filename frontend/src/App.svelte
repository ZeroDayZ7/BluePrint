<script lang="ts">
  import MainLayout from "./components/MainLayout.svelte";
  import Sidebar from "./components/Sidebar.svelte";
  import Console from "./components/Console.svelte";
  import Dashboard from "./routes/Dashboard.svelte";
  import { screens } from "./lib/router";
  import DeviceSelector from "./components/DeviceSelector.svelte";

  let activeTab = $state("dashboard");
  let deviceStatus = $state("Disconnected");
  let deviceName = $state("None");
  let logs = $state(["System initialized...", "Vite + Svelte 5 running..."]);

  const CurrentPage = $derived(screens[activeTab] || Dashboard);

  async function refreshDevices() {
    logs = [...logs, "Scanning for ADB devices..."];
  }
</script>

<MainLayout>
  {#snippet sidebar()}
    <Sidebar bind:activeTab {deviceStatus} />
  {/snippet}

  {#snippet header()}
    <DeviceSelector {deviceName} {deviceStatus} onRefresh={refreshDevices} />
  {/snippet}

  {#snippet content()}
    <div class="flex flex-col h-full">
      <div class="flex-1 overflow-y-auto custom-scrollbar">
        <div class="w-full flex justify-center p-8">
          <div class="w-full max-w-4xl">
            <CurrentPage bind:logs />
          </div>
        </div>
      </div>

      <Console {logs} />
    </div>
  {/snippet}
</MainLayout>
