<script lang="ts">
  let { value, class: className = "" } = $props();
</script>

<span
  class="inline-flex items-center justify-center px-1.5 py-0.5 rounded text-xs font-mono text-slate-300 min-w-[24px] {className}"
>
  {value}
</span>
