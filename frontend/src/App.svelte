<script lang="ts">
  import MainLayout from "./components/MainLayout.svelte";
  import Sidebar from "./components/Sidebar.svelte";
  import Console from "./components/Console.svelte";
  import Dashboard from "./routes/Dashboard.svelte";
  import { screens } from "./lib/router";
  import DeviceSelector from "./components/DeviceSelector.svelte";
  // import Titlebar from "./components/trash/Titlebar.svelte";

  let activeTab = $state("dashboard");
  let logs = $state(["System initialized...", "Vite + Svelte 5 running..."]);

  const CurrentPage = $derived(screens[activeTab] || Dashboard);
</script>

<!-- <Titlebar title="ADB COMMANDER | BLUE PRINT" /> -->

<MainLayout>
  {#snippet sidebar()}
    <Sidebar bind:activeTab />
  {/snippet}

  {#snippet header()}
    <DeviceSelector />
  {/snippet}

  {#snippet content()}
    <CurrentPage bind:logs />
  {/snippet}
</MainLayout>
