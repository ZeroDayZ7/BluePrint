<div
  class="flex flex-col h-full font-mono text-sm bg-black/40 rounded-xl border border-slate-800 overflow-hidden"
>
  <div class="p-4 flex-1 text-emerald-500 overflow-y-auto space-y-1">
    <div class="flex gap-2">
      <span class="text-blue-400">blueprint@android:~$</span>
      <span class="text-slate-200">adb shell</span>
    </div>
    <div class="text-slate-500 italic">
      # Terminal initialized. Ready for commands...
    </div>
  </div>

  <div
    class="p-3 bg-slate-900/80 border-t border-slate-800 flex items-center gap-3"
  >
    <span class="text-blue-400 font-bold">$</span>
    <input
      type="text"
      placeholder="Type adb shell command..."
      class="bg-transparent border-none outline-none text-slate-200 flex-1"
    />
  </div>
</div>
