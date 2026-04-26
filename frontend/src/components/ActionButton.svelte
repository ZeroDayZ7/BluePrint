<script lang="ts">
  interface Props {
    onclick: (e: MouseEvent) => void;
    label?: string;
    icon?: "trash" | "zap" | "download";
    variant?: "default" | "danger" | "success" | "warning";
    title?: string;
    disabled?: boolean;
  }

  let {
    onclick,
    label = "",
    icon = "zap",
    variant = "default",
    title = "",
    disabled = false,
  }: Props = $props();

  const variantClasses = {
    default: "text-slate-500 hover:text-blue-400 hover:bg-blue-500/10",
    danger: "text-slate-500 hover:text-red-400 hover:bg-red-500/10",
    success: "text-slate-500 hover:text-emerald-400 hover:bg-emerald-500/10",
    warning: "text-slate-500 hover:text-amber-400 hover:bg-amber-500/10",
  };
</script>

<button
  {onclick}
  {title}
  {disabled}
  type="button"
  aria-label={title || label || icon}
  class="
    opacity-0 group-hover:opacity-100
    flex items-center gap-1.5 px-2 py-1
    rounded-md transition-all duration-200
    cursor-pointer border-none bg-transparent
    active:scale-90 disabled:opacity-30 disabled:pointer-events-none
    {variantClasses[variant]}
  "
>
  <div class="flex items-center justify-center">
    {#if icon === "zap"}
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="12"
        height="12"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2.5"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <path d="M13 2 3 14h9l-1 8 10-12h-9l1-8z" />
      </svg>
    {:else if icon === "download"}
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="12"
        height="12"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2.5"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
        <polyline points="7 10 12 15 17 10" />
        <line x1="12" y1="15" x2="12" y2="3" />
      </svg>
    {:else if icon === "trash"}
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="12"
        height="12"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2.5"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <path d="M3 6h18" />
        <path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6" />
        <path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2" />
      </svg>
    {/if}
  </div>

  {#if label}
    <span class="text-[10px] font-bold uppercase tracking-tight leading-none">
      {label}
    </span>
  {/if}
</button>

<style>
  /* Dodatkowy delikatny efekt poświaty przy hoverze dla wariantów kolorystycznych */
  button:hover {
    text-shadow: 0 0 8px currentColor;
  }
</style>
