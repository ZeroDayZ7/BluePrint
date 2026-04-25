<script lang="ts">
  interface Props {
    deviceName: string;
    deviceStatus: string;
    onRefresh: () => Promise<void>;
  }

  let { deviceName, deviceStatus, onRefresh }: Props = $props();
  let isRefreshing = $state(false);

  const isConnected = $derived(deviceStatus === "Connected");

  async function handleRefresh() {
    if (isRefreshing) return;

    isRefreshing = true;

    // Timer na minimum 1 sekundę dla płynności animacji
    const minTime = new Promise((resolve) => setTimeout(resolve, 1000));

    try {
      // Czekamy na skanowanie i na upłynięcie sekundy
      await Promise.all([onRefresh(), minTime]);
    } finally {
      isRefreshing = false;
    }
  }
</script>

<div class="flex items-center gap-6">
  <div class="flex flex-col">
    <span
      class="text-[10px] text-slate-500 uppercase tracking-[0.2em] font-bold leading-none mb-1"
    >
      Active Device
    </span>
    <div class="flex items-center gap-2">
      <div
        class="w-2 h-2 rounded-full transition-colors {isConnected
          ? 'bg-blue-400 shadow-[0_0_8px_rgba(96,165,250,0.5)]'
          : 'bg-slate-600'}"
      ></div>
      <span
        class="text-sm font-mono {isConnected
          ? 'text-blue-400'
          : 'text-slate-400'}"
      >
        {deviceName}
      </span>
    </div>
  </div>

  <button
    onclick={handleRefresh}
    disabled={isRefreshing}
    class="p-2 rounded-lg bg-slate-800/50 border border-slate-700/50 hover:bg-slate-700 hover:border-slate-600 transition-all group active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed"
    title="Refresh ADB devices"
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="2"
      stroke="currentColor"
      class="w-4 h-4 {isRefreshing ? 'animate-spin' : ''}"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99"
      />
    </svg>
  </button>
</div>
