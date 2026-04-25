<script lang="ts">
  interface Props {
    value: string;
    placeholder?: string;
  }
  let { value = $bindable(), placeholder = "Search..." }: Props = $props();
</script>

<div class="relative flex-1">
  <div class="absolute left-3 top-1/2 -translate-y-1/2 text-slate-600">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="14"
      height="14"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      stroke-width="2"
      stroke-linecap="round"
      stroke-linejoin="round"
      ><circle cx="11" cy="11" r="8" /><path d="m21 21-4.3-4.3" /></svg
    >
  </div>
  <input
    type="text"
    bind:value
    {placeholder}
    class="w-full bg-slate-900/50 border border-slate-800 rounded-xl py-2 pl-10 pr-4 text-xs text-slate-300 focus:outline-none focus:border-blue-500/50 transition-all"
  />
</div>
