<script lang="ts">
  import Button from "./Button.svelte";

  interface Props {
    currentPath: string;
    storagePoints: string[];
    onNavigate: (path: string) => void;
    onBack: () => void;
  }

  let { currentPath, storagePoints, onNavigate, onBack }: Props = $props();

  let isAtRoot = $derived(storagePoints.includes(currentPath));

  // Pomocnicza funkcja do cięcia ścieżki, żeby nie robić tego w inline onclick
  function getPathUntil(parts: string[], index: number) {
    return "/" + parts.slice(0, index + 1).join("/");
  }
</script>

<div
  class="flex items-center rounded-md gap-2 px-2 h-10 border border-slate-800 bg-slate-900/60 mb-1"
>
  <div class="flex items-center">
    <Button
      variant="ghost"
      onclick={onBack}
      title="Go back"
      disabled={isAtRoot}
      class="!p-1.5 !rounded-md flex items-center justify-center disabled:opacity-20"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="14"
        height="14"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="3"
      >
        <path d="M11 17l-5-5 5-5M18 17l-5-5 5-5" />
      </svg>
    </Button>
  </div>

  <div
    class="flex items-center overflow-hidden overflow-x-auto no-scrollbar text-[12px] leading-none font-mono flex-1 min-w-0 h-full"
  >
    <button
      onclick={() => onNavigate(storagePoints[0] || "/sdcard")}
      class="text-blue-400 hover:bg-blue-500/10 px-1 rounded transition-colors shrink-0"
    >
      device
    </button>

    {#each currentPath.split("/").filter(Boolean) as segment, i (segment + i)}
      {@const allParts = currentPath.split("/").filter(Boolean)}
      <span class="flex items-center min-w-0 h-full">
        <span class="text-slate-600 shrink-0 px-1">/</span>
        <button
          onclick={() => onNavigate(getPathUntil(allParts, i))}
          class="text-slate-300 hover:text-blue-400 hover:bg-slate-800 px-1 rounded transition-colors truncate"
        >
          {segment}
        </button>
      </span>
    {/each}
  </div>
</div>
