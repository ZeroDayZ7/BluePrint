export const ADB_STATUS = {
  DEVICE: "device",
  OFFLINE: "offline",
  UNAUTHORIZED: "unauthorized",
  RECOVERY: "recovery",
  SIDELOAD: "sideload",
} as const;

export type AdbStatus = (typeof ADB_STATUS)[keyof typeof ADB_STATUS];
