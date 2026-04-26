<script lang="ts">
  import type { Snippet } from "svelte";

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
  class="flex flex-col p-4 bg-[#161b22]/40 rounded-2xl border border-slate-800/60 h-full"
>
  <div class="flex items-center justify-between mb-4">
    <div class="flex flex-col">
      <h3 class="text-xs font-black uppercase tracking-widest text-slate-500">
        {title}
      </h3>
      <span
        class="text-[10px] text-slate-600 font-medium truncate max-w-[200px]"
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

      <button
        onclick={onRefresh}
        type="button"
        title="Refresh"
        class="p-1.5 rounded-lg bg-slate-800/40 text-slate-400 hover:text-white hover:bg-slate-700 transition-all border border-slate-700/50 active:scale-95"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="16"
          height="16"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <path d="M21 12a9 9 0 1 1-9-9c2.52 0 4.93 1 6.74 2.74L21 8" />
          <path d="M21 3v5h-5" />
        </svg>
      </button>
    </div>
  </div>

  <div class="flex-1 min-h-0">
    {@render children()}
  </div>
</div>
