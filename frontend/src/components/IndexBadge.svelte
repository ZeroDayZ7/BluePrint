<script>
  let { value, class: className = "" } = $props();
</script>

<span class="text-xs {className}">
  {value}
</span>
