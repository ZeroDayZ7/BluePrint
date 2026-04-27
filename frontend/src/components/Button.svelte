<script lang="ts">
  import type { Snippet } from "svelte";
  import type { HTMLButtonAttributes } from "svelte/elements";

  interface Props extends HTMLButtonAttributes {
    children?: Snippet;
    variant?:
      | "primary"
      | "secondary"
      | "danger"
      | "ghost"
      | "action"
      | "actionDanger";
    size?: "sm" | "md" | "lg" | "icon";
    class?: string;
  }

  let {
    children,
    onclick,
    variant = "primary",
    size = "md",
    class: className = "",
    ...rest
  }: Props = $props();

  const baseClasses =
    "inline-flex items-center justify-center gap-2 font-medium transition-all duration-200 active:scale-95 disabled:opacity-30 disabled:pointer-events-none";

  const variants = {
    primary:
      "bg-blue-600 hover:bg-blue-500 text-white shadow-[0_0_15px_rgba(37,99,235,0.3)] rounded-xl",
    secondary:
      "bg-slate-800 hover:bg-slate-700 text-slate-200 border border-slate-700 rounded-xl",
    danger:
      "bg-red-600/10 hover:bg-red-600 text-red-500 hover:text-white border border-red-600/20 rounded-xl",
    ghost: "hover:bg-slate-800 text-slate-400 hover:text-white rounded-lg",
    action:
      "text-slate-500 hover:text-blue-400 hover:bg-blue-500/10 rounded-md opacity-0 group-hover:opacity-100",
    actionDanger:
      "text-slate-500 hover:text-red-400 hover:bg-red-500/10 rounded-md opacity-0 group-hover:opacity-100",
  };

  const sizes = {
    sm: "px-2 py-1 text-[10px] uppercase tracking-tight",
    md: "px-4 py-2 text-sm",
    lg: "px-6 py-3 text-base",
    icon: "p-2",
  };
</script>

<button
  {onclick}
  class="{baseClasses} {variants[variant]} {sizes[size]} {className}"
  {...rest}
>
  {#if children}
    {@render children()}
  {/if}
</button>

<style>
  button:not(:disabled):hover {
    text-shadow: 0 0 8px currentColor;
  }
</style>
