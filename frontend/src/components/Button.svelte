<script lang="ts">
  import type { Snippet } from "svelte";

  interface Props {
    children: Snippet;
    onclick?: () => void;
    variant?: "primary" | "secondary" | "danger" | "ghost";
    class?: string;
    [key: string]: any;
  }

  let {
    children,
    onclick,
    variant = "primary",
    class: className = "",
    ...rest
  }: Props = $props();

  const variants = {
    primary:
      "bg-blue-600 hover:bg-blue-500 text-white shadow-[0_0_15px_rgba(37,99,235,0.3)]",
    secondary:
      "bg-slate-800 hover:bg-slate-700 text-slate-200 border border-slate-700",
    danger:
      "bg-red-600/10 hover:bg-red-600 text-red-500 hover:text-white border border-red-600/20",
    ghost: "hover:bg-slate-800 text-slate-400 hover:text-white",
  };
</script>

<button
  {onclick}
  class="px-4 py-2 rounded-xl font-medium transition-all active:scale-95 flex items-center justify-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed {variants[
    variant
  ]} {className}"
  {...rest}
>
  {@render children()}
</button>
