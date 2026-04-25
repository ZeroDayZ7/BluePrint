<script lang="ts">
  import Button from "../components/Button.svelte";

  let { logs = $bindable() } = $props(); // Odbieramy logi z App.svelte
</script>

<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
  <div
    class="col-span-2 bg-[#161b22] border border-slate-800 rounded-2xl p-6 shadow-sm"
  >
    <h3 class="text-lg font-semibold mb-4">Quick Actions</h3>
    <div class="grid grid-cols-2 sm:grid-cols-4 gap-4">
      <Button variant="secondary" class="h-24 flex-col text-xs">
        Reboot Bootloader
      </Button>
      <Button variant="secondary" class="h-24 flex-col text-xs">
        Screen Mirror
      </Button>
      <Button variant="secondary" class="h-24 flex-col text-xs">
        Install APK
      </Button>
      <Button variant="danger" class="h-24 flex-col text-xs">
        Kill Server
      </Button>
    </div>
  </div>

  <div
    class="bg-[#05070a] border border-slate-800 rounded-2xl p-4 font-mono text-[11px] h-64 flex flex-col shadow-inner"
  >
    <div
      class="text-slate-500 mb-2 border-b border-slate-800 pb-1 font-bold uppercase tracking-widest text-[9px]"
    >
      ADB Terminal
    </div>
    <div class="flex-1 overflow-y-auto space-y-1 custom-scrollbar">
      {#each logs as log}
        <div class="flex gap-2">
          <span class="text-blue-600 font-bold">></span>
          <span class="text-slate-300">{log}</span>
        </div>
      {/each}
    </div>
  </div>
</div>
