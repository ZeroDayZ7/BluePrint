<script lang="ts">
  import type { Snippet } from "svelte";
  import Button from "./Button.svelte";
  import { RotateCw } from "lucide-svelte";

  interface Props {
    title: string;
    subtitle: string;
    searchPlaceholder?: string;
    searchQuery: string;
    onRefresh: () => void;
    headerExtra?: Snippet;
    searchActions?: Snippet;
    children: Snippet;
  }

  let {
    title,
    subtitle,
    searchPlaceholder = "Search...",
    searchQuery = $bindable(),
    onRefresh,
    headerExtra = undefined,
    searchActions = undefined,
    children,
  }: Props = $props();
</script>

<div
  class="flex flex-col p-4 bg-[#161b22]/40 rounded-xl border border-slate-800/60 h-full"
>
  <div class="flex items-center justify-between mb-4">
    <div class="flex flex-col">
      <h3 class="text-xs font-black uppercase tracking-widest text-slate-400">
        {title}
      </h3>
      <span
        class="text-[10px] text-slate-400 font-medium truncate max-w-[200px]"
      >
        {subtitle}
      </span>
    </div>

    {#if headerExtra}
      {@render headerExtra()}
    {/if}
  </div>

  <div class="flex items-center gap-2 mb-4">
    <div class="flex-1">
      <input
        type="text"
        bind:value={searchQuery}
        placeholder={searchPlaceholder}
        class="w-full bg-slate-900/50 border border-slate-800 rounded-lg px-3 py-1.5 text-[11px] text-slate-300 focus:outline-none focus:border-blue-500/50 transition-all"
      />
    </div>

    <div class="flex items-center gap-1">
      {#if searchActions}
        {@render searchActions()}
      {/if}

      <Button
        onclick={onRefresh}
        variant="secondary"
        size="icon"
        title="Refresh"
      >
        <RotateCw size={16} strokeWidth={2} />
      </Button>
    </div>
  </div>

  <div class="flex-1 min-h-0">
    {@render children()}
  </div>
</div>
