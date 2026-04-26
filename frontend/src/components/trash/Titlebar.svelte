<script lang="ts">
  import {
    WindowMinimise,
    WindowToggleMaximise,
    Quit,
  } from "../../../wailsjs/runtime/runtime";

  let { title = "ADB Commander Blue Print" } = $props();
</script>

<nav id="wails-titlebar">
  <div class="drag-region">
    <div class="logo">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="18"
        height="18"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="icon-pulse"
        ><path d="M12 2v20" /><path d="m4.93 4.93 14.14 14.14" /><path
          d="M2 12h20"
        /><path d="m4.93 19.07 14.14-14.14" /></svg
      >
      <span class="title">{title}</span>
    </div>
  </div>

  <div class="controls">
    <button onclick={WindowMinimise} aria-label="minimize">
      <svg width="10" height="1" viewBox="0 0 10 1" fill="none"
        ><rect width="10" height="1" fill="currentColor" /></svg
      >
    </button>
    <button onclick={WindowToggleMaximise} aria-label="maximize">
      <svg width="10" height="10" viewBox="0 0 10 10" fill="none"
        ><rect
          x="0.5"
          y="0.5"
          width="9"
          height="9"
          stroke="currentColor"
        /></svg
      >
    </button>
    <button onclick={Quit} class="close" aria-label="close">
      <svg width="10" height="10" viewBox="0 0 10 10" fill="none"
        ><path
          d="M1 1L9 9M9 1L1 9"
          stroke="currentColor"
          stroke-width="1.2"
        /></svg
      >
    </button>
  </div>
</nav>

<style>
  #wails-titlebar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 38px;
    background: rgba(15, 23, 42, 0.6);
    backdrop-filter: blur(10px);
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);
    color: #cbd5e1;
    user-select: none;
    z-index: 9999;
  }

  .drag-region {
    flex: 1;
    height: 100%;
    display: flex;
    align-items: center;
    padding-left: 12px;
    /* Flaga z main.go: CSSDragProperty */
    --wails-draggable: drag;
  }

  .logo {
    display: flex;
    align-items: center;
    gap: 10px;
    font-size: 13px;
    font-weight: 500;
    letter-spacing: 0.5px;
    color: #38bdf8; /* Cyan-blue akcent */
  }

  .icon-pulse {
    filter: drop-shadow(0 0 5px rgba(56, 189, 248, 0.5));
  }

  .controls {
    display: flex;
    height: 100%;
    --wails-draggable: none;
  }

  button {
    width: 46px;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    background: transparent;
    border: none;
    color: #94a3b8;
    transition: all 0.2s;
    cursor: pointer;
  }

  button:hover {
    background: rgba(255, 255, 255, 0.1);
    color: white;
  }

  button.close:hover {
    background: #ef4444;
    color: white;
  }

  .title {
    color: #f8fafc;
    opacity: 0.8;
  }
</style>
