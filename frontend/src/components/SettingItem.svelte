<script lang="ts">
  import Switch from "./Switch.svelte";

  interface Props {
    title: string;
    description: string;
    checked: boolean;
  }

  let { title, description, checked = $bindable() }: Props = $props();
</script>

<div class="flex items-start justify-between py-3 group w-full">
  <div class="flex flex-col text-left">
    <span
      class="text-sm font-medium text-slate-200 tracking-wide group-hover:text-blue-400 transition-colors italic"
    >
      {title}
    </span>
    <span class="text-xs text-slate-500 mt-0.5 leading-relaxed max-w-md">
      {description}
    </span>
  </div>

  <div class="pt-1">
    <Switch bind:checked />
  </div>
</div>
