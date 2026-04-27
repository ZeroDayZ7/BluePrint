<script lang="ts">
  import Button from "./Button.svelte";

  interface Tab {
    id: string;
    label: string;
    icon?: any;
  }

  interface Props {
    tabs: Tab[];
    activeTab: string;
    onChange: (id: string) => void;
  }

  let { tabs, activeTab, onChange }: Props = $props();
</script>

<div
  class="flex items-center gap-1 bg-slate-950 p-0.5 rounded-lg border border-slate-800/50 w-full"
>
  {#each tabs as tab}
    <Button
      onclick={() => onChange(tab.id)}
      variant={activeTab === tab.id ? "primary" : "ghost"}
      size="none"
      class="flex-1 py-1 text-xs font-bold tracking-blueprint uppercase transition-all"
    >
      {tab.label}
    </Button>
  {/each}
</div>
