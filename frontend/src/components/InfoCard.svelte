<script lang="ts">
  import type { HTMLAttributes } from "svelte/elements";

  interface Props extends HTMLAttributes<HTMLDivElement> {
    label: string;
    value: string | number;
    subValue?: string;
    variant?: "default" | "blue" | "green" | "amber";
    icon?: any;
  }

  let {
    label,
    value,
    subValue,
    variant = "default",
    icon: Icon,
    class: className,
    ...rest
  }: Props = $props();

  const styles = $derived(
    {
      default: "bg-slate-400/10 text-slate-400",
      blue: "bg-blue-500/20 text-blue-400",
      green: "bg-emerald-500/20 text-emerald-400",
      amber: "bg-amber-500/20 text-amber-400",
    }[variant],
  );
</script>

<div
  {...rest}
  class="{className ??
    ''}group flex items-center gap-4 p-3 bg-slate-900/40 border border-slate-800/50 rounded-xl"
>
  {#if Icon}
    <div
      class="flex items-center justify-center w-10 h-10 rounded-xl shrink-0 {styles}"
    >
      <Icon size={20} strokeWidth={2} />
    </div>
  {/if}

  <div class="flex flex-col min-w-0">
    <span
      class="text-[10px] uppercase font-bold tracking-[0.1em] text-slate-500 mb-0.5"
    >
      {label}
    </span>
    <div class="flex items-baseline gap-2">
      <span class="text-sm font-mono font-bold text-slate-200 truncate">
        {value}
      </span>
      {#if subValue}
        <span
          class="text-xs font-medium text-slate-500 bg-slate-800 px-1.5 py-0.5 rounded uppercase"
        >
          {subValue}
        </span>
      {/if}
    </div>
  </div>
</div>
