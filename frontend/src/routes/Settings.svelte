<script lang="ts">
  import Switch from "../components/Switch.svelte";

  let neonEffects = $state(true);
</script>

<div class="max-w-2xl bg-[#161b22] border border-slate-800 rounded-2xl p-8">
  <h2 class="text-2xl font-bold mb-6 tracking-tight text-white font-sans">
    Settings
  </h2>
  <div class="space-y-6">
    <div
      class="flex items-center justify-between p-4 bg-slate-900/40 rounded-xl border border-slate-800/50"
    >
      <div>
        <div class="font-medium text-slate-100">Dark Mode Neon</div>
        <div class="text-sm text-slate-500 font-sans">
          Glow effects on buttons and cards
        </div>
      </div>
      <Switch bind:checked={neonEffects} label="" />
    </div>

    <div class="flex flex-col gap-2 font-sans">
      <label for="adb-path" class="text-sm font-semibold text-slate-400 ml-1"
        >ADB Binary Path</label
      >
      <input
        id="adb-path"
        type="text"
        value="C:/platform-tools/adb.exe"
        class="bg-[#0b0e14] border border-slate-800 rounded-xl px-4 py-3 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500/50 text-slate-300"
      />
    </div>
  </div>
</div>
