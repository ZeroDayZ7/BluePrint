<script lang="ts">
  interface Props {
    checked: boolean;
    label?: string;
    onchange?: (val: boolean) => void;
    id?: string;
  }

  let {
    checked = $bindable(false),
    label = "",
    onchange,
    id = `switch-${Math.random().toString(36).slice(2, 9)}`,
  }: Props = $props();

  function toggle() {
    checked = !checked;
    onchange?.(checked);
  }
</script>

<div class="flex items-center gap-3">
  {#if label}
    <label
      for={id}
      class="text-[10px] uppercase font-semibold tracking-[0.2em] text-slate-500 cursor-pointer select-none"
    >
      {label}
    </label>
  {/if}

  <button
    type="button"
    {id}
    onclick={toggle}
    aria-label={label || "Toggle switch"}
    aria-pressed={checked}
    class="relative w-11 h-6 flex items-center rounded-full transition-colors duration-300
    {checked
      ? 'bg-blue-500/20 border-blue-400/40'
      : 'bg-slate-800/40 border-slate-700/40'}
    border backdrop-blur-md"
  >
    <!-- track highlight -->
    <span
      class="absolute inset-0 rounded-full transition-opacity duration-300
      {checked ? 'opacity-100 bg-blue-500/10' : 'opacity-0'}"
    ></span>

    <!-- knob -->
    <span
      class="relative h-4 w-4 rounded-full transition-transform duration-300 ease-out
      {checked
        ? 'translate-x-[22px] bg-blue-400'
        : 'translate-x-1 bg-slate-400'}"
    ></span>
  </button>
</div>
