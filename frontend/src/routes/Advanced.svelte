<script lang="ts">
  import LogcatView from "./advanced/LogcatView.svelte";
  import PowerMenu from "./advanced/PowerMenu.svelte";
  import TerminalView from "./advanced/TerminalView.svelte";
  import DeviceInfo from "./advanced/DeviceInfo.svelte";
  import SystemProps from "./advanced/SystemProps.svelte";
  import Tabs from "../components/Tabs.svelte";
  import type { Component } from "svelte";

  type TabId = "terminal" | "logcat" | "device" | "props" | "power";

  let activeTab = $state("terminal");

  const tabs: { id: TabId; label: string }[] = [
    { id: "terminal", label: "Terminal" },
    { id: "logcat", label: "Logcat" },
    { id: "device", label: "Device Info" },
    { id: "props", label: "System" },
    { id: "power", label: "Power Menu" },
  ];

  const views: Record<TabId, Component<any>> = {
    terminal: TerminalView,
    logcat: LogcatView,
    device: DeviceInfo,
    props: SystemProps,
    power: PowerMenu,
  };
</script>

<div class="flex flex-col gap-4 h-full overflow-hidden">
  <Tabs {tabs} {activeTab} onChange={(id) => (activeTab = id)} />
  <div
    class="flex-1 bg-slate-900/20 border border-slate-800/40 rounded-md relative"
  >
    {#if views[activeTab]}
      {@const ActiveView = views[activeTab]}
      <ActiveView />
    {/if}
  </div>
</div>
