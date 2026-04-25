<script lang="ts">
  import Button from "../components/Button.svelte";
  import { StartMirroring } from "../../wailsjs/go/main/App";
  import { deviceState } from "../lib/deviceState.svelte";

  let { logs = $bindable() } = $props();

  async function handleMirror() {
    if (!deviceState.activeDevice) {
      logs = [...logs, "Error: No device selected"];
      return;
    }

    try {
      const result = await StartMirroring(deviceState.activeDevice.id);
      logs = [
        ...logs,
        `Mirroring: ${result} for ${deviceState.activeDevice.id}`,
      ];
    } catch (err) {
      logs = [...logs, `Error: ${err}`];
    }
  }
</script>

<div class="space-y-8">
  <div class="bg-[#161b22]/40 border border-slate-800/60 rounded-2xl p-6">
    <h3
      class="text-sm font-black uppercase tracking-widest text-slate-500 mb-6"
    >
      Quick Actions
    </h3>
    <div class="grid grid-cols-2 sm:grid-cols-4 gap-4">
      <Button variant="secondary" class="h-14 flex-col text-xs">
        Reboot Bootloader
      </Button>

      <Button
        variant="secondary"
        class="h-14 flex-col text-xs"
        onclick={handleMirror}
        disabled={!deviceState.isConnected}
      >
        Screen Mirror
      </Button>

      <Button variant="secondary" class="h-14 flex-col text-xs">
        Install APK
      </Button>

      <Button variant="danger" class="h-14 flex-col text-xs text-red-400">
        Kill Server
      </Button>
    </div>
  </div>
</div>
