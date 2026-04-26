<script>
  let { value, class: className = "" } = $props();
</script>

<span class="text-[9px] font-bold text-slate-500 {className}">
  {value}
</span>
