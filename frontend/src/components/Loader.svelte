<script lang="ts">
  interface Props {
    message?: string;
    size?: "sm" | "md" | "lg";
  }

  let { message = "Loading...", size = "md" }: Props = $props();

  const sizes = {
    sm: "size-5 border-2",
    md: "size-8 border-[3px]",
    lg: "size-12 border-4",
  };
</script>

<div class="flex flex-col items-center justify-center p-8 gap-4">
  <div class="relative">
    <div class="{sizes[size]} rounded-full border-slate-800"></div>

    <div
      class="{sizes[size]} 
             absolute top-0 left-0
             rounded-full
             border-t-blue-500 border-r-transparent border-b-transparent border-l-transparent
             animate-spin
             shadow-[0_0_15px_rgba(59,130,246,0.5)]"
    ></div>
  </div>

  {#if message}
    <p
      class="text-xxs font-black uppercase tracking-blueprint text-slate-500 animate-pulse"
    >
      {message}
    </p>
  {/if}
</div>
