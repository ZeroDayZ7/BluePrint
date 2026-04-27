<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
  {#each ["System", "Recovery", "Bootloader", "Fastboot"] as mode}
    <div
      class="p-6 bg-slate-900/40 border border-slate-800 rounded-md flex items-center justify-between group hover:border-blue-500/30 transition-all"
    >
      <div>
        <h3 class="text-slate-200 font-bold">Reboot to {mode}</h3>
        <p class="text-[11px] text-slate-500 italic">
          Restart device into {mode.toLowerCase()} mode
        </p>
      </div>
      <button
        type="button"
        aria-label="Reboot to {mode}"
        title="Reboot to {mode}"
        class="p-3 bg-slate-800 rounded-md group-hover:bg-blue-600 group-hover:text-white transition-all"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="20"
          height="20"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <path
            d="M12 2v4M12 18v4M4.93 4.93l2.83 2.83M16.24 16.24l2.83 2.83M2 12h4M18 12h4M4.93 19.07l2.83-2.83M16.24 7.76l2.83-2.83"
          />
        </svg>
      </button>
    </div>
  {/each}
</div>
