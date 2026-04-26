<script lang="ts">
  interface Props {
    onclick: (e: MouseEvent) => void;
    label: string;
    icon?: "trash" | "zap" | "download"; // Dodałem download do typów
    variant?: "default" | "danger" | "success";
    title?: string;
  }

  let {
    onclick,
    label,
    icon = "trash",
    variant = "default",
    title = "",
  }: Props = $props();

  // Mapowanie kolorów na klasy Tailwind
  const variantClasses = {
    default: "text-slate-500 hover:text-blue-400",
    danger: "text-red-500/60 hover:text-red-400",
    success: "text-emerald-500/60 hover:text-emerald-400",
  };
</script>

<button
  {onclick}
  {title}
  class="opacity-0 group-hover:opacity-100 flex items-center gap-1.5 px-1 py-0.5
         transition-all duration-200 cursor-pointer border-none bg-transparent
         {variantClasses[variant]}"
>
  {#if icon === "zap"}
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="11"
      height="11"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      stroke-width="2.5"
      stroke-linecap="round"
      stroke-linejoin="round"><path d="M13 2 3 14h9l-1 8 10-12h-9l1-8z" /></svg
    >
  {:else if icon === "download"}
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="11"
      height="11"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      stroke-width="2.5"
      stroke-linecap="round"
      stroke-linejoin="round"
      ><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" /><polyline
        points="7 10 12 15 17 10"
      /><line x1="12" y1="15" x2="12" y2="3" /></svg
    >
  {:else}
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="11"
      height="11"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      stroke-width="2.5"
      stroke-linecap="round"
      stroke-linejoin="round"
      ><path d="M3 6h18" /><path
        d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6"
      /><path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2" /></svg
    >
  {/if}

  {#if label}
    <span class="text-[10px] font-black uppercase tracking-tighter">
      {label}
    </span>
  {/if}
</button>
