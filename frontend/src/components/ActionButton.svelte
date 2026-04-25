<script lang="ts">
  interface Props {
    onclick: () => void;
    label: string;
    icon?: "trash" | "zap";
  }
  let { onclick, label, icon = "trash" }: Props = $props();
</script>

<button
  {onclick}
  class="opacity-0 group-hover:opacity-100 flex items-center gap-1.5 px-1 py-0.5
         text-red-500/60 hover:text-red-400
         transition-all duration-200 cursor-pointer border-none bg-transparent"
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

  <span class="text-[10px] font-black uppercase tracking-tighter">
    {label}
  </span>
</button>
