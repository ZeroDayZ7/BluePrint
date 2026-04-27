<script lang="ts">
  interface Property {
    key: string;
    value: string;
  }

  interface Props {
    items: Property[];
    height?: string;
  }

  let { items, height = "h-[305px]" }: Props = $props();
</script>

<div
  class="{height} rounded-xl border border-slate-800/50 bg-[#0b0e14]/50 flex flex-col overflow-hidden"
>
  <table
    class="w-full text-left text-xs border-collapse bg-[#161b22] border-b border-slate-800 shrink-0"
  >
    <thead>
      <tr>
        <th
          class="px-4 py-2 font-bold text-slate-500 uppercase tracking-widest w-1/2"
          >Key</th
        >
        <th
          class="px-4 py-2 font-bold text-slate-500 uppercase tracking-widest w-1/2"
          >Value</th
        >
      </tr>
    </thead>
  </table>

  <div class="flex-1 overflow-y-auto custom-scrollbar">
    <table class="w-full text-left text-[11px] border-collapse table-fixed">
      <tbody class="divide-y divide-slate-800/30">
        {#each items as prop}
          <tr class="hover:bg-blue-500/5 transition-colors group">
            <td
              class="px-4 py-2 font-mono text-slate-500 group-hover:text-blue-400 break-all border-r border-slate-800/30 w-1/2"
            >
              {prop.key}
            </td>
            <td class="px-4 py-2 font-mono text-slate-300 break-all w-1/2">
              {prop.value}
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</div>
