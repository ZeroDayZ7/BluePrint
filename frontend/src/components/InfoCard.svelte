<script lang="ts">
  interface Props {
    label: string;
    value: string | number;
    subValue?: string;
    variant?: "default" | "blue" | "green" | "amber";
    icon?: any;
  }

  let {
    label,
    value,
    subValue = "",
    variant = "default",
    icon: Icon,
  }: Props = $props();

  const themes = {
    default: "bg-slate-800/40 border-slate-700/50 text-slate-400",
    blue: "bg-blue-500/10 border-blue-500/20 text-blue-400",
    green: "bg-emerald-500/10 border-emerald-500/20 text-emerald-400",
    amber: "bg-amber-500/10 border-amber-500/20 text-amber-400",
  };

  const textColors = {
    default: "text-slate-200",
    blue: "text-blue-100",
    green: "text-emerald-100",
    amber: "text-amber-100",
  };
</script>

<div
  class="group relative overflow-hidden p-3 border rounded-xl flex flex-col justify-center min-h-[70px] transition-colors duration-200 {themes[
    variant
  ]}"
>
  <div
    class="absolute right-3 top-1/2 -translate-y-1/2 opacity-10 pointer-events-none"
  >
    {#if Icon}
      <Icon size={42} strokeWidth={1.5} />
    {/if}
  </div>

  <div class="relative z-10 flex items-center gap-2 mb-1.5">
    {#if Icon}
      <Icon size={13} strokeWidth={2.5} class="shrink-0" />
    {/if}
    <span class="text-[10px] uppercase font-black tracking-widest opacity-70">
      {label}
    </span>
  </div>

  <div class="relative z-10 flex items-baseline gap-2">
    <span
      class="text-l uppercase font-mono leading-none tracking-tight {textColors[
        variant
      ]}"
    >
      {value}
    </span>
    {#if subValue}
      <span class="text-[10px] font-semibold opacity-40 uppercase">
        {subValue}
      </span>
    {/if}
  </div>
</div>
