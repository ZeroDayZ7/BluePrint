import { GetSettings, SaveSettings } from "../../wailsjs/go/main/App";
import type { config } from "../../wailsjs/go/models";

class ConfigStoreManager {
  public data = $state<config.AppConfig>({
    showHidden: false,
    alwaysTop: false,
    devMode: false,
    autoConnect: false,
    adbPath: "C:/platform-tools/adb.exe",
  });

  private isInitialized = false;

  constructor() {
    this.loadInitialConfig();
  }

  private async loadInitialConfig() {
    try {
      const savedConfig = await GetSettings();

      this.data = {
        ...this.data,
        ...Object.fromEntries(
          Object.entries(savedConfig).filter(
            ([_, v]) => v !== "" && v !== null,
          ),
        ),
      };

      $effect.root(() => {
        $effect(() => {
          if (this.isInitialized) {
            this.persist();
          }
        });
      });

      this.isInitialized = true;
    } catch (err) {
      console.error("Failed to load config:", err);
    }
  }

  public async persist() {
    await SaveSettings(this.data);
  }
}

export const configStore = new ConfigStoreManager();
