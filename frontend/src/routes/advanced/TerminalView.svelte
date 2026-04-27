<script lang="ts">
  import { ExecuteShell } from "../../../wailsjs/go/main/App";
  import { deviceState } from "../../lib/deviceState.svelte";
  import { tick } from "svelte";

  // Definiujemy typ wpisu, aby TS nie zgadywał
  interface TerminalLine {
    type: "cmd" | "out";
    text: string;
  }

  let command = $state("");
  let history = $state<TerminalLine[]>([]); // Jawne przypisanie typu
  let inputElement = $state<HTMLInputElement | null>(null);
  let isLoading = $state(false);

  const MAX_HISTORY_LINES = 500;

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

      // Używamy "as const", aby TS wiedział, że "cmd" to nie jest zwykły string
      const cmdEntry: TerminalLine = { type: "cmd", text: currentCmd };
      history = [...history, cmdEntry].slice(-MAX_HISTORY_LINES);

      command = "";
      isLoading = true;

      try {
        const result = await ExecuteShell(deviceId, currentCmd);
        const outEntry: TerminalLine = {
          type: "out",
          text: result || "(no output)",
        };
        history = [...history, outEntry].slice(-MAX_HISTORY_LINES);
      } catch (err) {
        const errEntry: TerminalLine = { type: "out", text: "Error: " + err };
        history = [...history, errEntry].slice(-MAX_HISTORY_LINES);
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
  class="flex flex-col h-[500px] font-mono text-sm bg-black/60 rounded-md border border-slate-800 overflow-hidden shadow-2xl"
>
  <div
    class="p-4 flex-1 overflow-y-auto flex flex-col custom-scrollbar selection:bg-blue-500/40 selection:text-white"
  >
    <div
      class="text-slate-500 text-xxs italic mb-4 border-b border-slate-800/30 pb-2 tracking-blueprint uppercase"
    >
      # Terminal ready for {deviceState.activeDevice?.model || "unknown"} (Buffer:
      {MAX_HISTORY_LINES})
    </div>

    {#each history as line}
      {#if line.type === "cmd"}
        <div class="flex gap-2 mt-3 first:mt-0">
          <span class="text-blue-400 shrink-0 font-bold"
            >blueprint@android:~$</span
          >
          <span class="text-slate-200 break-all">{line.text}</span>
        </div>
      {:else}
        <div
          class="text-emerald-500 whitespace-pre-wrap break-all leading-relaxed mt-1 mb-2 font-light"
        >
          {line.text}
        </div>
      {/if}
    {/each}

    {#if isLoading}
      <div class="text-blue-400 animate-pulse mt-2 flex items-center gap-2">
        <span class="inline-block w-2 h-4 bg-blue-400"></span>
        <span class="text-xxs uppercase tracking-blueprint font-bold"
          >Executing...</span
        >
      </div>
    {/if}
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
        ? "Enter shell command..."
        : "Connect device..."}
      class="bg-transparent border-none outline-none text-slate-200 flex-1 disabled:opacity-50 placeholder:text-slate-600"
    />
  </div>
</div>
