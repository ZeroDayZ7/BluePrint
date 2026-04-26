<script lang="ts">
  import { deviceState } from "../lib/deviceState.svelte";
  import ListContainer from "../components/ListContainer.svelte";
  import IndexBadge from "../components/IndexBadge.svelte";
  import ActionButton from "../components/ActionButton.svelte";
  import { untrack } from "svelte";

  let isLoading = $state(false);
  let searchQuery = $state("");

  // Reaktywne pobieranie plików z cache lub backendu
  let files = $derived(
    deviceState.activeDevice
      ? deviceState.filesCache[deviceState.activeDevice.id]?.[
          deviceState.currentPath
        ] || []
      : [],
  );

  let filteredFiles = $derived(
    files.filter((f) =>
      f.name.toLowerCase().includes(searchQuery.toLowerCase()),
    ),
  );

  async function loadDirectory(path: string, force = false) {
    const device = deviceState.activeDevice;
    if (!device || isLoading) return;

    if (!force && deviceState.getCachedFiles(device.id, path)) {
      deviceState.currentPath = path;
      return;
    }

    isLoading = true;
    try {
      // @ts-ignore - Wywołanie wygenerowanego backendu Wails
      const result = await window.go.main.App.ListFiles(device.id, path);

      untrack(() => {
        if (!deviceState.filesCache[device.id])
          deviceState.filesCache[device.id] = {};
        deviceState.filesCache[device.id][path] = result;
        deviceState.currentPath = path;
      });
    } catch (err) {
      console.error("FS Error:", err);
    } finally {
      isLoading = false;
    }
  }

  function navigateUp() {
    const parts = deviceState.currentPath.split("/").filter(Boolean);
    parts.pop();
    loadDirectory("/" + parts.join("/"));
  }

  // Funkcja pomocnicza do nawigacji klawiaturą (a11y)
  function handleKeydown(e: KeyboardEvent, file: any) {
    if (e.key === "Enter" && file.isDir) {
      loadDirectory(
        `${deviceState.currentPath}/${file.name}`.replace(/\/+/g, "/"),
      );
    }
  }

  $effect(() => {
    if (deviceState.isConnected && deviceState.activeDevice) {
      untrack(() => loadDirectory(deviceState.currentPath));
    }
  });
</script>

<ListContainer
  title="File Explorer"
  subtitle={deviceState.currentPath}
  searchPlaceholder="Search files..."
  bind:searchQuery
  onRefresh={() => loadDirectory(deviceState.currentPath, true)}
>
  {#snippet headerExtra()}
    <div class="flex items-center gap-2">
      <button
        onclick={navigateUp}
        title="Go back"
        aria-label="Navigate up"
        class="p-1.5 hover:bg-slate-800 rounded-md text-slate-400 transition-colors"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="16"
          height="16"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"><path d="M11 17l-5-5 5-5M18 17l-5-5 5-5" /></svg
        >
      </button>
      <div
        class="bg-slate-900/50 px-3 py-1 rounded-md border border-slate-800 text-[11px] font-mono text-slate-400"
      >
        {deviceState.currentPath}
      </div>
    </div>
  {/snippet}

  {#snippet searchActions()}
    <div class="w-2"></div>
  {/snippet}

  <div class="h-[310px] overflow-y-auto custom-scrollbar">
    <div
      class="grid grid-cols-12 px-3 py-2 text-[10px] font-bold text-slate-500 uppercase tracking-tighter border-b border-slate-800/30 mb-2"
    >
      <div class="col-span-7">Name</div>
      <div class="col-span-2 text-center">Type</div>
      <div class="col-span-3 text-right">Actions</div>
    </div>

    {#if isLoading}
      <div
        class="flex h-32 items-center justify-center text-slate-500 text-xs italic"
      >
        Accessing filesystem...
      </div>
    {:else}
      <div class="space-y-0.5">
        {#each filteredFiles as file}
          <div
            role="button"
            tabindex="0"
            onclick={() =>
              file.isDir &&
              loadDirectory(
                `${deviceState.currentPath}/${file.name}`.replace(/\/+/g, "/"),
              )}
            onkeydown={(e) => handleKeydown(e, file)}
            class="grid grid-cols-12 items-center px-3 py-2 hover:bg-blue-500/5 rounded-lg group cursor-pointer transition-all border border-transparent hover:border-blue-500/10"
          >
            <div class="col-span-7 flex items-center gap-3">
              {#if file.isDir}
                <svg
                  class="text-blue-500"
                  xmlns="http://www.w3.org/2000/svg"
                  width="16"
                  height="16"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  ><path
                    d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 2h9a2 2 0 0 1 2 2z"
                  ></path></svg
                >
              {:else}
                <svg
                  class="text-slate-500"
                  xmlns="http://www.w3.org/2000/svg"
                  width="16"
                  height="16"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  ><path
                    d="M13 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V9z"
                  ></path><polyline points="13 2 13 9 20 9"></polyline></svg
                >
              {/if}
              <span
                class="text-xs {file.isDir
                  ? 'text-slate-200'
                  : 'text-slate-400'} truncate font-medium"
              >
                {file.name}
              </span>
            </div>

            <div class="col-span-2 text-center">
              <IndexBadge value={file.isDir ? "DIR" : "FILE"} />
            </div>

            <div class="col-span-3 flex justify-end gap-1">
              <ActionButton
                icon="zap"
                label=""
                title="Download"
                onclick={() => {
                  /* Pull file */
                }}
              />
              <ActionButton
                icon="trash"
                variant="danger"
                label=""
                title="Delete"
                onclick={() => {
                  /* Delete */
                }}
              />
            </div>
          </div>
        {/each}
      </div>
    {/if}
  </div>
</ListContainer>
