<script lang="ts">
  let { title, description } = $props();
</script>

<header class="flex items-baseline justify-between pb-6 mb-8">
  <h2 class="text-2xl font-bold tracking-tight text-white font-sans">
    {title}
  </h2>
  <p class="text-sm text-slate-500 font-medium italic">
    {description}
  </p>
</header>
