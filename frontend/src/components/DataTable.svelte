<script lang="ts" generics="T">
  import type { Snippet } from "svelte";

  interface Props {
    items: T[];
    height?: string;
    header: Snippet;
    row: Snippet<[T, number]>;
  }

  let {
    items,
    height = "h-[calc(100vh-288px)]",
    header,
    row,
  }: Props = $props();
</script>

<div
  class="{height} rounded-md border border-slate-800/50 bg-brand-surface/50 flex flex-col overflow-hidden shadow-2xl"
>
  <div
    class="shrink-0 border-b border-slate-800 bg-slate-900/80 backdrop-blur-sm"
  >
    {@render header()}
  </div>

  <div class="flex-1 overflow-y-auto custom-scrollbar bg-slate-950/20">
    <div class="divide-y divide-slate-800/20">
      {#each items as item, i}
        <div class="hover:bg-blue-500/5 transition-colors group">
          {@render row(item, i)}
        </div>
      {/each}
    </div>
  </div>
</div>
