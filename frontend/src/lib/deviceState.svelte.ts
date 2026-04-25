// src/lib/deviceState.svelte.ts
interface Device {
  id: string;
  model: string;
  status: string;
}

class DeviceState {
  devices = $state<Device[]>([]);
  selectedDeviceIndex = $state(0);
  isRefreshing = $state(false);

  activeDevice = $derived(
    this.devices.length > 0 ? this.devices[this.selectedDeviceIndex] : null,
  );

  displayStatus = $derived.by(() => {
    const device = this.activeDevice;
    if (!device) return "Disconnected";

    if (device.status === "device" || device.status === "Connected") {
      return "Connected";
    }
    if (device.status === "unauthorized") return "Unauthorized";
    if (device.status === "offline") return "Offline";

    return device.status;
  });

  isConnected = $derived(this.displayStatus === "Connected");
}

export const deviceState = new DeviceState();
