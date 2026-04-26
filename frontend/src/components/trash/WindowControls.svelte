<script lang="ts">
  import {
    Quit,
    WindowMinimise,
    WindowToggleMaximise,
  } from "../../../wailsjs/runtime/runtime";
</script>

<div class="window-controls" style="--wails-draggable: none;">
  <button onclick={WindowMinimise} class="control-btn" title="Minimalizuj">
    <svg width="12" height="12" viewBox="0 0 12 12"
      ><rect x="2" y="5.5" width="8" height="1" fill="currentColor" /></svg
    >
  </button>

  <button
    onclick={WindowToggleMaximise}
    class="control-btn"
    title="Maksymalizuj"
  >
    <svg width="12" height="12" viewBox="0 0 12 12"
      ><rect
        x="2.5"
        y="2.5"
        width="7"
        height="7"
        stroke="currentColor"
        fill="none"
        stroke-width="1"
      /></svg
    >
  </button>

  <button onclick={Quit} class="control-btn close-btn" title="Zamknij">
    <svg width="12" height="12" viewBox="0 0 12 12"
      ><path
        d="M2.5 2.5L9.5 9.5M9.5 2.5L2.5 9.5"
        stroke="currentColor"
        stroke-width="1.2"
      /></svg
    >
  </button>
</div>

<style>
  .window-controls {
    display: flex;
    gap: 4px;
    align-items: center;
  }

  .control-btn {
    width: 32px;
    height: 32px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 6px;
    border: none;
    background: transparent;
    color: #94a3b8; /* Slate-400 */
    transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
    cursor: pointer;
    padding: 0;
  }

  .control-btn:hover {
    background: rgba(255, 255, 255, 0.08);
    color: #f1f5f9; /* Slate-100 */
  }

  .close-btn:hover {
    background: #ef4444 !important; /* Red-500 */
    color: white !important;
  }

  .control-btn:active {
    transform: scale(0.92);
  }
</style>
