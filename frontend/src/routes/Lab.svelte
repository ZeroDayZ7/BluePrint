<script lang="ts">
  import { Zap } from "lucide-svelte";
  import Button from "../components/Button.svelte";
  import ListContainer from "../components/ListContainer.svelte";

  let searchQuery = $state("");
  let testValue = $state("Prototype Mode Active");

  function testAlert() {
    alert("Lab test successful! Backend connection ready.");
  }
</script>

<div class="space-y-4 h-full">
  <ListContainer
    title="Experimental Lab"
    subtitle="Test your new features here"
    bind:searchQuery
    onRefresh={() => console.log("Refreshed Lab")}
    searchPlaceholder="Search lab items..."
  >
    <div
      class="p-8 flex flex-col items-center justify-center border-2 border-dashed border-slate-800 rounded-md bg-slate-900/20"
    >
      <div class="text-blue-500 mb-4">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="48"
          height="48"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="1"
          stroke-linecap="round"
          stroke-linejoin="round"
          ><path
            d="M10 2v8L4.72 20.55a1 1 0 0 0 .9 1.45h12.76a1 1 0 0 0 .9-1.45L14 10V2"
          /><path d="M8.5 2h7" /><path d="M7 16h10" /></svg
        >
      </div>

      <h2 class="text-slate-300 font-bold mb-2 text-lg">
        Welcome to the Sandbox
      </h2>
      <p class="text-slate-500 text-sm mb-6 text-center max-w-md">
        To jest bezpieczna strefa. Możesz tu wrzucać nowe przyciski, testować
        parsery lub układy gridów bez wpływu na resztę aplikacji.
      </p>

      <div class="flex gap-4">
        <Button variant="primary" onclick={testAlert}>
          <Zap
            size={14}
            strokeWidth={2.5}
            fill="currentColor"
            class="text-amber-400"
          />
          <span>Trigger Test</span>
        </Button>
        <div
          class="px-4 py-2 bg-slate-800/50 rounded-lg text-[11px] font-mono text-blue-400 border border-slate-700"
        >
          Status: {testValue}
        </div>
      </div>
    </div>

    <div class="mt-6 grid grid-cols-2 gap-4">
      <div class="p-4 bg-slate-900/40 border border-slate-800 rounded-lg">
        <h4 class="text-[10px] font-black uppercase text-slate-600 mb-2">
          UI Component Test
        </h4>
        <input type="range" class="w-full accent-blue-500" />
      </div>
      <div class="p-4 bg-slate-900/40 border border-slate-800 rounded-lg">
        <h4 class="text-[10px] font-black uppercase text-slate-600 mb-2">
          Input Echo
        </h4>
        <div class="text-xs text-slate-400 font-mono italic">
          "{searchQuery}"
        </div>
      </div>
    </div>
  </ListContainer>
</div>
