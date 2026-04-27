<script lang="ts">
  import { deviceState } from "../lib/deviceState.svelte";
  import ListContainer from "../components/ListContainer.svelte";
  import IndexBadge from "../components/IndexBadge.svelte";
  import { untrack } from "svelte";
  import Button from "../components/Button.svelte";
  import Breadcrumbs from "../components/Breadcrumbs.svelte";
  import { Download, Trash2 } from "lucide-svelte";

  let isLoading = $state(false);
  let showHidden = $state(false);
  let searchQuery = $state("");
  let storagePoints = $state<string[]>(["/sdcard"]);
  let pendingDelete = $state<string | null>(null);

  interface FileItem {
    name: string;
    path: string;
    isDir: boolean;
    size: string;
    modifiedAt?: string;
  }

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

  $effect(() => {
    const deviceId = deviceState.activeDevice?.id;
    if (deviceState.isConnected && deviceId) {
      // @ts-ignore
      window.go.main.App.GetStoragePoints(deviceState.activeDevice.id).then(
        (points) => {
          storagePoints = points;
        },
      );
    }
  });

  async function changeStorage(newPath: string) {
    if (!deviceState.activeDevice) return;

    deviceState.currentPath = newPath;

    if (!deviceState.getCachedFiles(deviceState.activeDevice.id, newPath)) {
      await loadDirectory(newPath, true);
    }
  }

  $effect(() => {
    if (showHidden !== undefined) {
      untrack(() => loadDirectory(deviceState.currentPath, true));
    }
  });

  async function loadDirectory(path: string, force = false) {
    const device = deviceState.activeDevice;
    if (!device || isLoading) return;

    if (!force && deviceState.getCachedFiles(device.id, path)) {
      deviceState.currentPath = path;
      return;
    }

    isLoading = true;
    try {
      // @ts-ignore
      const result = await window.go.main.App.ListFiles(
        device.id,
        path,
        showHidden,
      );

      untrack(() => {
        if (!deviceState.filesCache[device.id])
          deviceState.filesCache[device.id] = {};
        deviceState.filesCache[device.id][path] = result;
        deviceState.currentPath = path;
      });
    } catch (err) {
      console.error(err);
    } finally {
      isLoading = false;
    }
  }

  async function downloadFile(file: FileItem) {
    const fullPath = `${deviceState.currentPath}/${file.name}`.replace(
      /\/+/g,
      "/",
    );
    // @ts-ignore
    await window.go.main.App.DownloadFile(
      deviceState.activeDevice!.id,
      fullPath,
      file.name,
    );
  }

  async function handleDelete(file: FileItem) {
    if (pendingDelete !== file.name) {
      pendingDelete = file.name;
      setTimeout(() => {
        if (pendingDelete === file.name) pendingDelete = null;
      }, 3000);
      return;
    }

    const fullPath = `${deviceState.currentPath}/${file.name}`.replace(
      /\/+/g,
      "/",
    );
    // @ts-ignore
    const success = await window.go.main.App.DeleteFile(
      deviceState.activeDevice!.id,
      fullPath,
    );

    if (success) {
      pendingDelete = null;
      loadDirectory(deviceState.currentPath, true);
    }
  }

  function navigateUp() {
    const parts = deviceState.currentPath.split("/").filter(Boolean);
    parts.pop();
    loadDirectory("/" + parts.join("/"));
  }

  function handleKeydown(e: KeyboardEvent, file: FileItem) {
    if (e.key === "Enter" && file.isDir) {
      loadDirectory(
        `${deviceState.currentPath}/${file.name}`.replace(/\/+/g, "/"),
      );
    }
  }

  $effect(() => {
    const deviceId = deviceState.activeDevice?.id;
    if (deviceState.isConnected && deviceId) {
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
    <div class="flex items-center gap-4">
      <label class="flex items-center gap-2 cursor-pointer group select-none">
        <div class="relative">
          <input
            type="checkbox"
            bind:checked={showHidden}
            class="sr-only peer"
          />
          <div
            class="w-7 h-4 bg-slate-800 rounded-full peer peer-checked:bg-blue-600 transition-colors after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-slate-400 after:rounded-full after:h-3 after:w-3 after:transition-all peer-checked:after:translate-x-3 peer-checked:after:bg-white"
          ></div>
        </div>
        <span
          class="text-[10px] font-bold text-slate-500 group-hover:text-slate-300 uppercase tracking-widest transition-colors"
        >
          Hidden
        </span>
      </label>
    </div>
  {/snippet}

  <Breadcrumbs
    currentPath={deviceState.currentPath}
    {storagePoints}
    onNavigate={(path) => loadDirectory(path)}
    onBack={navigateUp}
  />

  {#snippet searchActions()}
    <div class="flex items-center">
      <select
        value={storagePoints.find((p) =>
          deviceState.currentPath.startsWith(p),
        ) || "/sdcard"}
        onchange={(e) => changeStorage(e.currentTarget.value)}
        class="bg-slate-800 border border-slate-700 text-xs text-slate-300 rounded px-2 py-1.5 outline-none focus:border-blue-500/50 transition-all cursor-pointer hover:bg-slate-700 appearance-none"
      >
        {#each storagePoints as point}
          <option value={point} class="bg-slate-800 text-slate-300 py-1">
            {point === "/sdcard"
              ? "📱 Internal Storage"
              : `💾 SD Card (${point.split("/").pop()})`}
          </option>
        {/each}
      </select>
    </div>
  {/snippet}

  <div class="overflow-y-auto custom-scrollbar">
    <div
      class="grid grid-cols-12 px-3 py-2 text-[10px] font-bold text-slate-500 uppercase tracking-tighter border-b border-slate-800/30 mb-2"
    >
      <div class="col-span-7">Name</div>
      <div class="col-span-2 text-center">Type</div>
      <div class="col-span-3 text-right text-red-500/70 italic">
        {pendingDelete ? "Click again to confirm delete" : "Actions"}
      </div>
    </div>

    {#if isLoading}
      <div
        class="flex h-32 items-center justify-center text-slate-500 text-xs italic"
      >
        Accessing filesystem...
      </div>
    {:else}
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
            <Button
              variant="action"
              size="icon"
              title="Download"
              disabled={pendingDelete !== null}
              onclick={(e: MouseEvent) => {
                e.stopPropagation();
                downloadFile(file);
              }}
            >
              <Download size={14} strokeWidth={2.5} />
            </Button>

            <Button
              variant={pendingDelete === file.name ? "danger" : "actionDanger"}
              size={pendingDelete === file.name ? "sm" : "icon"}
              title={pendingDelete === file.name ? "Confirm Delete" : "Delete"}
              onclick={(e: MouseEvent) => {
                e.stopPropagation();
                handleDelete(file);
              }}
            >
              <Trash2 size={14} strokeWidth={2.5} />
              {#if pendingDelete === file.name}
                <span class="ml-1">SURE?</span>
              {/if}
            </Button>
          </div>
        </div>
      {/each}

      {#if filteredFiles.length === 0}
        <div
          class="p-8 text-center text-slate-600 text-xs uppercase tracking-widest italic"
        >
          No files found
        </div>
      {/if}
    {/if}
  </div>
</ListContainer>
