<script lang="ts">
  import MainLayout from "./components/MainLayout.svelte";
  import Sidebar from "./components/Sidebar.svelte";
  import Button from "./components/Button.svelte";
  import Dashboard from "./routes/Dashboard.svelte";
  import { screens } from "./lib/router";

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
    <div class="text-xs font-medium text-slate-500 uppercase tracking-widest">
      Device: <span class="text-blue-400 ml-1 font-mono">{deviceName}</span>
    </div>
    <Button onclick={refreshDevices} class="text-xs px-5 py-2">
      Refresh ADB
    </Button>
  {/snippet}

  {#snippet content()}
    <CurrentPage bind:logs />
  {/snippet}
</MainLayout>
