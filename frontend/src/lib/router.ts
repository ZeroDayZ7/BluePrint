// BluePrint\frontend\src\lib\router.ts
import Dashboard from "../routes/Dashboard.svelte";
import Settings from "../routes/Settings.svelte";
import Files from "../routes/Files.svelte";
import Applications from "../routes/Applications.svelte";

export const navigation = [
  { id: "dashboard", label: "Dashboard", icon: "home" },
  { id: "apps", label: "Applications", icon: "grid" },
  { id: "files", label: "File Explorer", icon: "folder" },
  { id: "settings", label: "Settings", icon: "settings" },
];

export const screens: Record<string, any> = {
  dashboard: Dashboard,
  apps: Applications,
  settings: Settings,
  files: Files,
};
