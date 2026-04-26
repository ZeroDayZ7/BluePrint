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

interface DeviceInfo {
  model: string;
  manufacturer: string;
  androidVer: string;
  apiLevel: string;
  cpu: string;
  battery: string;
  serial: string;
}

class DeviceState {
  devices = $state<Device[]>([]);
  selectedDeviceIndex = $state(0);
  isRefreshing = $state(false);

  deviceInfoCache = $state<Record<string, DeviceInfo>>({});
  systemPropsCache = $state<Record<string, { key: string; value: string }[]>>(
    {},
  );
  userAppsCache = $state<Record<string, string[]>>({});
  systemAppsCache = $state<Record<string, string[]>>({});
  processesCache = $state<Record<string, any[]>>({});
  filesCache = $state<Record<string, Record<string, FileEntry[]>>>({});
  currentPath = $state("/sdcard");

  clearDeviceCache(deviceId: string) {
    delete this.deviceInfoCache[deviceId];
    delete this.systemPropsCache[deviceId];
    delete this.userAppsCache[deviceId];
    delete this.systemAppsCache[deviceId];
  }

  getCachedFiles(deviceId: string, path: string) {
    return this.filesCache[deviceId]?.[path] || null;
  }

  clearPathCache(deviceId: string, path: string) {
    if (this.filesCache[deviceId]) {
      delete this.filesCache[deviceId][path];
    }
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
