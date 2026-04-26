<script lang="ts">
  import Button from "./Button.svelte";

  interface Tab {
    id: string;
    label: string;
    icon?: string;
  }

  interface Props {
    tabs: Tab[];
    activeTab: string;
    onChange: (id: string) => void;
  }

  let { tabs, activeTab, onChange }: Props = $props();
</script>

<div
  class="flex items-center gap-1 bg-slate-900/50 p-1 rounded-xl border border-slate-800/50 w-full"
>
  {#each tabs as tab}
    <Button
      onclick={() => onChange(tab.id)}
      variant={activeTab === tab.id ? "primary" : "ghost"}
      class="flex-1 !rounded-lg !py-2 !text-sm"
    >
      {tab.label}
    </Button>
  {/each}
</div>
