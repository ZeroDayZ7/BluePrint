<script lang="ts">
  import { Minus, Square, X } from "lucide-svelte";
  import {
    Quit,
    WindowMinimise,
    WindowToggleMaximise,
  } from "../../../wailsjs/runtime/runtime";
  import Button from "../Button.svelte";
</script>

<div class="flex items-center h-full" style="--wails-draggable: none;">
  <Button
    variant="system"
    size="none"
    class="w-[46px] h-full"
    onclick={WindowMinimise}
    title="Minimize"
  >
    <Minus size={14} />
  </Button>

  <Button
    variant="system"
    size="none"
    class="w-[46px] h-full"
    onclick={WindowToggleMaximise}
    title="Maximize"
  >
    <Square size={10} />
  </Button>

  <Button
    variant="systemDanger"
    size="none"
    class="w-[46px] h-full"
    onclick={Quit}
    title="Close"
  >
    <X size={16} />
  </Button>
</div>
