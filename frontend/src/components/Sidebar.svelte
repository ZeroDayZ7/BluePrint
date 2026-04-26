<script lang="ts">
  import Logo from "./Logo.svelte";
  import StatusBadge from "./StatusBadge.svelte";
  import { navigation } from "../lib/router";

  interface Props {
    activeTab: string;
  }

  let { activeTab = $bindable() }: Props = $props();
</script>

<nav class="w-64 bg-[#0d1117] border-r border-slate-800/60 flex flex-col p-5">
  <div class="mb-10">
    <Logo />
  </div>

  <div class="space-y-1.5 flex-1">
    {#each navigation as tab}
      <button
        onclick={() => (activeTab = tab.id)}
        class="w-full flex items-center gap-3 px-4 py-3 rounded-xl transition-all capitalize font-medium
        {activeTab === tab.id
          ? 'bg-blue-600/10 text-blue-400 border border-blue-500/20 shadow-sm'
          : 'text-slate-500 hover:text-slate-300 hover:bg-slate-800/40'}"
      >
        {tab.label}
      </button>
    {/each}
  </div>

  <StatusBadge />
</nav>
