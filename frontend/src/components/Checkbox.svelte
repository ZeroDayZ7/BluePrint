<script lang="ts">
  interface Props {
    checked: boolean;
    onchange: (val: boolean) => void;
  }
  let { checked = $bindable(), onchange }: Props = $props();
</script>

<button
  type="button"
  onclick={() => {
    checked = !checked;
    onchange?.(checked);
  }}
  class="w-4 h-4 rounded border transition-all flex items-center justify-center
  {checked
    ? 'bg-blue-600 border-blue-500 shadow-[0_0_10px_rgba(37,99,235,0.4)]'
    : 'bg-slate-900 border-slate-700 hover:border-slate-500'}"
>
  {#if checked}
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="12"
      height="12"
      viewBox="0 0 24 24"
      fill="none"
      stroke="white"
      stroke-width="4"
      stroke-linecap="round"
      stroke-linejoin="round"><polyline points="20 6 9 17 4 12" /></svg
    >
  {/if}
</button>
