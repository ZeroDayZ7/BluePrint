<div class="flex flex-col h-full gap-4">
  <div class="flex justify-between items-center px-2">
    <div class="text-[10px] uppercase tracking-widest text-slate-500 font-bold">
      Live Stream
    </div>
    <div class="flex gap-2">
      <button
        type="button"
        class="px-3 py-1 bg-red-500/10 text-red-500 rounded-md text-xs hover:bg-red-500/20"
      >
        Clear
      </button>
      <button
        type="button"
        class="px-3 py-1 bg-blue-500 text-white rounded-md text-xs"
      >
        Pause
      </button>
    </div>
  </div>

  <div
    class="flex-1 bg-slate-950 rounded-xl border border-slate-800 p-3 font-mono text-[11px] overflow-y-auto space-y-1"
  >
    <div class="text-blue-400">
      [04-26 13:29:01.442] I/ActivityManager: START u0 {"{"}act =
      android.intent.action.MAIN{"}"}
    </div>

    <div class="text-emerald-400">
      [04-26 13:29:01.510] D/BluetoothAdapter: getState() : STATE_ON
    </div>

    <div class="text-amber-400">
      [04-26 13:29:02.102] W/InputMethodManager: Ignoring onBindInput
    </div>

    <div class="text-red-400">
      [04-26 13:29:03.001] E/AndroidRuntime: FATAL EXCEPTION: main
    </div>
  </div>
</div>
