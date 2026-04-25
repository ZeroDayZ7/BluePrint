// BluePrint\frontend\src\lib\router.ts
import Dashboard from "../routes/Dashboard.svelte";
import Settings from "../routes/Settings.svelte";
import Files from "../routes/Files.svelte";

export const navigation = [
  { id: "dashboard", label: "Dashboard", icon: "home" },
  { id: "files", label: "File Explorer", icon: "folder" },
  { id: "settings", label: "Settings", icon: "settings" },
];

export const screens: Record<string, any> = {
  dashboard: Dashboard,
  settings: Settings,
  files: Files,
};
