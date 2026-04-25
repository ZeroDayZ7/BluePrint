<script lang="ts">
  import Button from "../components/Button.svelte";
  import Switch from "../components/Switch.svelte"; // Zakładam że go tam użyjesz

  let { logs = $bindable() } = $props();
</script>

<div class="space-y-8">
  <div class="bg-[#161b22]/40 border border-slate-800/60 rounded-2xl p-6">
    <h3
      class="text-sm font-black uppercase tracking-widest text-slate-500 mb-6"
    >
      Quick Actions
    </h3>
    <div class="grid grid-cols-2 sm:grid-cols-4 gap-4">
      <Button variant="secondary" class="h-24 flex-col text-xs"
        >Reboot Bootloader</Button
      >
      <Button variant="secondary" class="h-24 flex-col text-xs"
        >Screen Mirror</Button
      >
      <Button variant="secondary" class="h-24 flex-col text-xs"
        >Install APK</Button
      >
      <Button variant="danger" class="h-24 flex-col text-xs text-red-400"
        >Kill Server</Button
      >
    </div>
  </div>

  <div class="px-2 space-y-6">
    <h3 class="text-sm font-black uppercase tracking-widest text-slate-500">
      Global Settings
    </h3>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-12">
      <div class="flex items-center justify-between group">
        <div>
          <div class="text-slate-200 font-medium tracking-wide">
            Developer Mode Monitoring
          </div>
          <p class="text-xs text-slate-500 mt-1">
            Keep track of device state changes in real-time.
          </p>
        </div>
        <Switch checked={true} />
      </div>

      <div class="flex items-center justify-between group">
        <div>
          <div class="text-slate-200 font-medium tracking-wide">
            Auto-Connect USB
          </div>
          <p class="text-xs text-slate-500 mt-1">
            Automatically authorize new devices via ADB.
          </p>
        </div>
        <Switch checked={false} />
      </div>
    </div>
  </div>
</div>
