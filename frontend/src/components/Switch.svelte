<script lang="ts">
  interface Props {
    checked: boolean;
    label?: string;
    onchange?: (val: boolean) => void;
    id?: string;
  }

  // Bindable pozwala na dwukierunkową synchronizację (bind:checked)
  let {
    checked = $bindable(false),
    label = "",
    onchange,
    id = `switch-${Math.random().toString(36).slice(2, 9)}`,
  }: Props = $props();

  function handleChange(e: Event) {
    const target = e.target as HTMLInputElement;
    checked = target.checked;
    onchange?.(checked);
  }
</script>

<div class="flex items-center gap-3">
  {#if label}
    <label
      for={id}
      class="text-sm font-medium text-slate-400 cursor-pointer select-none"
    >
      {label}
    </label>
  {/if}

  <div class="relative inline-flex items-center">
    <input
      type="checkbox"
      {id}
      class="sr-only peer"
      {checked}
      onchange={handleChange}
    />

    <button
      type="button"
      {id}
      onclick={() => (checked = !checked)}
      class="w-11 h-6 bg-slate-800 rounded-full border border-slate-700
             peer-checked:bg-blue-600 peer-checked:border-blue-500
             peer-focus:ring-2 peer-focus:ring-blue-500/50
             transition-all duration-300 cursor-pointer
             shadow-[inset_0_2px_4px_rgba(0,0,0,0.2)]"
      aria-pressed={checked}
      aria-label={label || "Toggle switch"}
    >
      <span
        class="absolute top-[3px] left-[3px] w-4 h-4 bg-white rounded-full
               transition-all duration-300 shadow-md
               {checked ? 'translate-x-5' : 'translate-x-0'}"
      ></span>
    </button>
  </div>
</div>

<style>
</style>
