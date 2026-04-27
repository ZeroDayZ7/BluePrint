<script lang="ts">
  import PageHeader from "../components/PageHeader.svelte";
  import SettingGroup from "../components/SettingGroup.svelte";
  import SettingItem from "../components/SettingItem.svelte";

  // Stan ustawień
  let devMode = $state(true);
  let autoConnect = $state(false);
  let adbPath = $state("C:/platform-tools/adb.exe");
</script>

<div class="max-w-3xl">
  <PageHeader
    title="Settings"
    description="Configure BluePrint and ADB environment"
  />

  <div class="space-y-2">
    <SettingGroup title="Global Settings" />

    <div class="flex flex-col">
      <SettingItem
        title="Developer Mode Monitoring"
        description="Keep track of device state changes in real-time."
        bind:checked={devMode}
      />

      <div class="h-[1px] bg-slate-800/30 w-full my-1"></div>

      <SettingItem
        title="Auto-Connect USB"
        description="Automatically authorize new devices via ADB."
        bind:checked={autoConnect}
      />
    </div>

    <SettingGroup title="ADB Configuration" />

    <div class="flex flex-col gap-4 mt-6">
      <div class="flex flex-col gap-1.5">
        <label
          for="adb-path"
          class="text-[10px] uppercase font-black tracking-widest text-slate-500 ml-1"
        >
          ADB Binary Path
        </label>
        <p class="text-xs text-slate-600 ml-1">
          Point to your adb.exe if it's not in system PATH.
        </p>
      </div>

      <div class="flex gap-3">
        <input
          id="adb-path"
          type="text"
          bind:value={adbPath}
          class="flex-1 bg-[#0b0e14] border border-slate-800 rounded-md px-4 py-3 text-sm outline-none text-slate-300 transition-all font-mono"
        />
        <button
          class="px-6 py-2 bg-slate-800/60 hover:bg-slate-700 border border-slate-700/50 rounded-md text-xs font-bold transition-all active:scale-95"
        >
          Browse
        </button>
      </div>
    </div>
  </div>
</div>
