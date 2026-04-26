<script lang="ts">
  import type { Snippet } from "svelte";

  interface Props {
    sidebar: Snippet;
    header: Snippet;
    content: Snippet;
  }

  let { sidebar, header, content }: Props = $props();
</script>

<main
  class="flex h-screen bg-[#0b0e14] text-slate-200 overflow-hidden select-none font-sans"
>
  {@render sidebar()}

  <section class="flex-1 flex flex-col">
    <header
      class="h-16 border-b border-slate-800/60 flex items-center justify-between px-8 bg-[#0b0e14]/80 backdrop-blur-xl z-10"
    >
      {@render header()}
    </header>

    <div class="flex-1 p-4 overflow-hidden">
      {@render content()}
    </div>
  </section>
</main>
