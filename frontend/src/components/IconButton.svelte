<script lang="ts">
  interface Props {
    onclick: () => void;
    variant?: "default" | "danger" | "success";
    title?: string;
    children?: any;
  }
  let { onclick, variant = "default", title, children }: Props = $props();

  const variants = {
    default: "text-slate-500 hover:bg-slate-800 hover:text-blue-400",
    danger: "text-slate-500 hover:bg-red-500/10 hover:text-red-400",
    success: "text-slate-500 hover:bg-green-500/10 hover:text-green-400",
  };
</script>

<button
  {onclick}
  {title}
  class="p-2 rounded-lg transition-all duration-200 {variants[variant]}"
>
  {@render children()}
</button>
