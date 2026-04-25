<script lang="ts">
  import Button from "../components/Button.svelte";
</script>

<div
  class="flex flex-col items-center justify-center h-full text-center space-y-6"
>
  <div
    class="w-20 h-20 bg-slate-800/50 rounded-3xl flex items-center justify-center border border-slate-700 shadow-xl"
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="w-10 h-10 text-blue-500"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z"
      />
    </svg>
  </div>

  <div class="max-w-xs">
    <h2 class="text-xl font-bold text-white mb-2 italic">File Explorer</h2>
    <p class="text-slate-500 text-sm">
      Browse and manage files on your Android device. This module is currently
      under development.
    </p>
  </div>

  <Button variant="secondary" class="text-xs px-8">Work in Progress</Button>
</div>
