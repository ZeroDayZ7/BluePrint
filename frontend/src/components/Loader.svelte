<script lang="ts">
  interface Props {
    message?: string;
    size?: "sm" | "md" | "lg";
  }

  let { message = "Loading...", size = "md" }: Props = $props();

  const sizes = {
    sm: "w-5 h-5 border-2",
    md: "w-8 h-8 border-3",
    lg: "w-12 h-12 border-4",
  };
</script>

<div class="flex flex-col items-center justify-center p-8 space-y-4">
  <div class="relative">
    <div class="{sizes[size]} rounded-full border-slate-800"></div>

    <div
      class="{sizes[
        size
      ]} rounded-full border-t-blue-500 border-r-transparent border-b-transparent border-l-transparent animate-spin absolute top-0 left-0 shadow-[0_0_15px_rgba(59,130,246,0.5)]"
    ></div>
  </div>

  {#if message}
    <p
      class="text-[10px] font-black uppercase tracking-[0.2em] text-slate-500 animate-pulse"
    >
      {message}
    </p>
  {/if}
</div>

<style>
  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }
  .animate-spin {
    animation: spin 0.8s linear infinite;
  }
</style>
