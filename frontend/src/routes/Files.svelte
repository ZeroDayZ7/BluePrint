<script lang="ts">
	import { deviceState } from "../lib/deviceState.svelte";
	import ListContainer from "../components/ListContainer.svelte";
	import IndexBadge from "../components/IndexBadge.svelte";
	import { untrack } from "svelte";
	import Button from "../components/Button.svelte";
	import Breadcrumbs from "../components/Breadcrumbs.svelte";
	import { Download, Folder, Trash2, File } from "lucide-svelte";
	import DataTable from "../components/DataTable.svelte";
	import Loader from "../components/Loader.svelte";

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
			? deviceState.filesCache[deviceState.activeDevice.id]?.[deviceState.currentPath] || []
			: [],
	);

	let filteredFiles = $derived(
		files.filter((f) => f.name.toLowerCase().includes(searchQuery.toLowerCase())),
	);

	$effect(() => {
		const deviceId = deviceState.activeDevice?.id;
		if (deviceState.isConnected && deviceId) {
			// @ts-ignore
			window.go.main.App.GetStoragePoints(deviceState.activeDevice.id).then((points) => {
				storagePoints = points;
			});
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
			const result = await window.go.main.App.ListFiles(device.id, path, showHidden);

			untrack(() => {
				if (!deviceState.filesCache[device.id]) deviceState.filesCache[device.id] = {};
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
		const fullPath = `${deviceState.currentPath}/${file.name}`.replace(/\/+/g, "/");
		// @ts-ignore
		await window.go.main.App.DownloadFile(deviceState.activeDevice!.id, fullPath, file.name);
	}

	async function handleDelete(file: FileItem) {
		if (pendingDelete !== file.name) {
			pendingDelete = file.name;
			setTimeout(() => {
				if (pendingDelete === file.name) pendingDelete = null;
			}, 3000);
			return;
		}

		const fullPath = `${deviceState.currentPath}/${file.name}`.replace(/\/+/g, "/");
		// @ts-ignore
		const success = await window.go.main.App.DeleteFile(deviceState.activeDevice!.id, fullPath);

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
			loadDirectory(`${deviceState.currentPath}/${file.name}`.replace(/\/+/g, "/"));
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
					<input type="checkbox" bind:checked={showHidden} class="sr-only peer" />
					<div
						class="w-7 h-4 bg-slate-800 rounded-full peer peer-checked:bg-blue-600 transition-colors after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-slate-400 after:rounded-full after:h-3 after:w-3 after:transition-all peer-checked:after:translate-x-3 peer-checked:after:bg-white"
					></div>
				</div>
				<span
					class="text-[10px] font-bold text-slate-500 group-hover:text-slate-300 uppercase tracking-widest transition-colors"
					>Hidden</span
				>
			</label>
		</div>
	{/snippet}

	<Breadcrumbs
		currentPath={deviceState.currentPath}
		class="mb-2"
		{storagePoints}
		onNavigate={(path) => loadDirectory(path)}
		onBack={navigateUp}
	/>

	{#snippet searchActions()}
		<div class="flex items-center">
			<select
				value={storagePoints.find((p) => deviceState.currentPath.startsWith(p)) || "/sdcard"}
				onchange={(e) => changeStorage(e.currentTarget.value)}
				class="bg-slate-800 border border-slate-700 text-xs text-slate-300 rounded px-2 py-1.5 outline-none focus:border-blue-500/50 transition-all cursor-pointer hover:bg-slate-700 appearance-none"
			>
				{#each storagePoints as point}
					<option value={point} class="bg-slate-800 text-slate-300 py-1">
						{point === "/sdcard" ? "📱 Internal Storage" : `💾 SD Card (${point.split("/").pop()})`}
					</option>
				{/each}
			</select>
		</div>
	{/snippet}

	{#if isLoading}
		<div class="flex h-64 items-center justify-center">
			<Loader message="Accessing filesystem..." size="md" />
		</div>
	{:else}
		<DataTable items={filteredFiles} height="min-h-50">
			{#snippet header()}
				<div
					class="grid grid-cols-12 px-4 py-2.5 text-[10px] font-black text-slate-500 uppercase tracking-widest"
				>
					<div class="col-span-7">Name</div>
					<div class="col-span-2 text-center">Type</div>
					<div
						class="col-span-3 text-right {pendingDelete ? 'text-red-500 italic animate-pulse' : ''}"
					>
						{pendingDelete ? "Confirm delete" : "Actions"}
					</div>
				</div>
			{/snippet}

			{#snippet row(file)}
				<div
					role="button"
					tabindex="0"
					onclick={() =>
						file.isDir &&
						loadDirectory(`${deviceState.currentPath}/${file.name}`.replace(/\/+/g, "/"))}
					onkeydown={(e) => handleKeydown(e, file)}
					class="grid grid-cols-12 items-center px-4 py-1.5 cursor-pointer"
				>
					<div class="col-span-7 flex items-center gap-3">
						{#if file.isDir}
							<Folder size={16} class="text-blue-500" strokeWidth={2} />
						{:else}
							<File size={16} class="text-slate-500" strokeWidth={2} />
						{/if}
						<span
							class="text-xs {file.isDir
								? 'text-slate-200 font-bold'
								: 'text-slate-400'} truncate font-mono"
						>
							{file.name}
						</span>
					</div>

					<div class="col-span-2 text-center">
						<IndexBadge value={file.isDir ? "DIR" : "FILE"} />
					</div>

					<div
						class="col-span-3 flex justify-end gap-1 opacity-0 group-hover:opacity-100 transition-opacity"
					>
						<Button
							variant="action"
							size="icon"
							disabled={pendingDelete !== null}
							onclick={(e: MouseEvent) => {
								e.stopPropagation();
								downloadFile(file);
							}}
						>
							<Download size={14} />
						</Button>

						<Button
							variant={pendingDelete === file.name ? "danger" : "actionDanger"}
							size={pendingDelete === file.name ? "sm" : "icon"}
							onclick={(e: MouseEvent) => {
								e.stopPropagation();
								handleDelete(file);
							}}
						>
							<Trash2 size={14} />
							{#if pendingDelete === file.name}
								<span class="ml-1 text-[10px] font-bold">SURE?</span>
							{/if}
						</Button>
					</div>
				</div>
			{/snippet}
		</DataTable>
	{/if}
</ListContainer>
