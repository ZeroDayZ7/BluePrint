<script lang="ts">
  import type { Snippet } from "svelte";

  interface Props {
    sidebar: Snippet;
    header: Snippet;
    content: Snippet;
  }

  let { sidebar, header, content }: Props = $props();
</script>

<main class="flex h-screen overflow-hidden font-sans">
  {@render sidebar()}

  <section class="flex flex flex-col w-full">
    <header class="border-b border-slate-800/60 p-2">
      {@render header()}
    </header>

    <div class="flex-1 p-4 overflow-hidden">
      {@render content()}
    </div>
  </section>
</main>
