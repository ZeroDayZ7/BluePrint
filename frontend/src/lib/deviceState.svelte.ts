// src/lib/deviceState.svelte.ts
interface Device {
  id: string;
  model: string;
  status: string;
}

interface FileEntry {
  name: string;
  size: string;
  date: string;
  time: string;
  isDir: boolean;
  path: string;
}

class DeviceState {
  devices = $state<Device[]>([]);
  selectedDeviceIndex = $state(0);
  isRefreshing = $state(false);

  userAppsCache = $state<Record<string, string[]>>({});
  systemAppsCache = $state<Record<string, string[]>>({});
  processesCache = $state<Record<string, any[]>>({});
  filesCache = $state<Record<string, Record<string, FileEntry[]>>>({});
  currentPath = $state("/sdcard");

  getCachedFiles(deviceId: string, path: string) {
    return this.filesCache[deviceId]?.[path] || null;
  }

  activeDevice = $derived(
    this.devices.length > 0 ? this.devices[this.selectedDeviceIndex] : null,
  );

  displayStatus = $derived.by(() => {
    const device = this.activeDevice;
    if (!device) return "Disconnected";
    if (device.status === "device" || device.status === "Connected")
      return "Connected";
    if (device.status === "unauthorized") return "Unauthorized";
    if (device.status === "offline") return "Offline";
    return device.status;
  });

  isConnected = $derived(this.displayStatus === "Connected");
}

export const deviceState = new DeviceState();
