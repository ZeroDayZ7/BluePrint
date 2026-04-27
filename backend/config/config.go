package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Definicja Twoich ustawień
type AppConfig struct {
	ShowHidden  bool   `json:"showHidden"`
	AlwaysTop   bool   `json:"alwaysTop"`
	DevMode     bool   `json:"devMode"`
	AutoConnect bool   `json:"autoConnect"`
	AdbPath     string `json:"adbPath"`
}

func GetConfigPath() string {
	appData := os.Getenv("LOCALAPPDATA")
	dir := filepath.Join(appData, "ADBCommander")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	return filepath.Join(dir, "config.json")
}

func Load() *AppConfig {
	cfg := &AppConfig{
		ShowHidden: false,
		AlwaysTop:  false,
	}

	data, err := os.ReadFile(GetConfigPath())
	if err != nil {
		Save(cfg)
		return cfg
	}

	json.Unmarshal(data, cfg)
	return cfg
}

func Save(cfg *AppConfig) error {
	data, _ := json.MarshalIndent(cfg, "", "  ")
	return os.WriteFile(GetConfigPath(), data, 0644)
}
