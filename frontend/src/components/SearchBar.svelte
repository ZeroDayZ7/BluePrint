<script lang="ts">
  import { Search } from "lucide-svelte";

  interface Props {
    value: string;
    placeholder?: string;
  }

  let { value = $bindable(), placeholder = "Search..." }: Props = $props();
</script>

<div class="relative flex-1 group">
  <div
    class="absolute left-3 top-1/2 -translate-y-1/2 text-slate-600 group-focus-within:text-blue-500 transition-colors"
  >
    <Search size={14} strokeWidth={2.5} />
  </div>
  <input
    type="text"
    bind:value
    {placeholder}
    class="w-full bg-slate-900/50 border border-slate-800 rounded-xl py-2 pl-9 pr-4 text-xxs text-slate-300 placeholder:text-slate-600 focus:outline-none focus:border-blue-500/50 focus:bg-slate-900 transition-all"
  />
</div>
