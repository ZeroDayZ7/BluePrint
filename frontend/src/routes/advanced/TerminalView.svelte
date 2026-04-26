<script lang="ts">
  import { ExecuteShell } from "../../../wailsjs/go/main/App";
  import { deviceState } from "../../lib/deviceState.svelte";
  import { tick } from "svelte";

  let command = $state("");
  let history = $state<{ type: "cmd" | "out"; text: string }[]>([]);
  let inputElement = $state<HTMLInputElement | null>(null);
  let isLoading = $state(false);

  $effect(() => {
    if (deviceState.isConnected && inputElement) {
      inputElement.focus();
    }
  });

  async function handleCommand(e: KeyboardEvent) {
    if (e.key === "Enter" && command.trim() !== "") {
      const currentCmd = command.trim();
      const deviceId = deviceState.activeDevice?.id;

      if (!deviceId) return;

      history = [{ type: "cmd", text: currentCmd }, ...history];
      command = "";
      isLoading = true;

      try {
        const result = await ExecuteShell(deviceId, currentCmd);
        history = [{ type: "out", text: result || "(no output)" }, ...history];
      } catch (err) {
        history = [{ type: "out", text: "Error: " + err }, ...history];
      } finally {
        isLoading = false;

        await tick();
        if (inputElement) {
          inputElement.focus();
        }
      }
    }
  }
</script>

<div
  class="flex flex-col h-[500px] font-mono text-sm bg-black/60 rounded-xl border border-slate-800 overflow-hidden shadow-2xl"
>
  <div
    class="p-4 flex-1 overflow-y-auto flex flex-col-reverse custom-scrollbar selection:bg-blue-500/40 selection:text-white"
  >
    <div>
      {#if isLoading}
        <div class="text-blue-400 animate-pulse mb-2">_ running command...</div>
      {/if}

      {#each [...history].reverse() as line}
        {#if line.type === "cmd"}
          <div class="flex gap-2 mt-3 first:mt-0">
            <span class="text-blue-400 shrink-0">blueprint@android:~$</span>
            <span class="text-slate-200 break-all">{line.text}</span>
          </div>
        {:else}
          <div
            class="text-emerald-500 whitespace-pre-wrap break-all leading-relaxed mt-1 mb-2"
          >
            {line.text}
          </div>
        {/if}
      {/each}

      <div class="text-slate-500 text-xs italic mb-4">
        # Terminal ready for {deviceState.activeDevice?.model || "unknown"}
      </div>
    </div>
  </div>

  <div
    class="p-3 bg-slate-900 border-t border-slate-800 flex items-center gap-3"
  >
    <span class="text-blue-400 font-bold shrink-0">$</span>
    <input
      bind:this={inputElement}
      type="text"
      bind:value={command}
      onkeydown={handleCommand}
      disabled={isLoading || !deviceState.isConnected}
      placeholder={deviceState.isConnected
        ? "Enter shell command (e.g. ls, getprop)..."
        : "Connect device..."}
      class="bg-transparent border-none outline-none text-slate-200 flex-1 disabled:opacity-50 placeholder:text-slate-600"
    />
  </div>
</div>
