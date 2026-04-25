package main

import (
	"BluePrint/backend/adb"
	"BluePrint/backend/mirror"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type App struct {
	ctx context.Context
}

type Device struct {
	ID     string `json:"id"`
	Model  string `json:"model"`
	Status string `json:"status"`
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) getADBPath() string {
	return "adb"
}

func (a *App) getToolPath(toolName string) string {
	cwd, _ := os.Getwd()
	if toolName == "adb" || toolName == "scrcpy" {
		toolName += ".exe"
	}
	return filepath.Join(cwd, "bin", "adb_tools", toolName)
}

func (a *App) GetDevices() []Device {
	adbPath := a.getToolPath("adb")
	rawOutput, err := adb.FetchDevices(fmt.Sprintf(`"%s" devices -l`, adbPath))

	if err != nil {
		return []Device{}
	}

	var devices []Device
	lines := strings.Split(rawOutput, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "List of devices") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) >= 2 {
			id := parts[0]
			status := parts[1]
			model := "Unknown"

			if strings.Contains(line, "model:") {
				mParts := strings.Split(line, "model:")
				if len(mParts) > 1 {
					model = strings.Fields(mParts[1])[0]
				}
			}

			devices = append(devices, Device{
				ID:     id,
				Status: status,
				Model:  model,
			})
		}
	}
	return devices
}

func (a *App) StartMirroring(deviceID string) string {
	adbPath := a.getToolPath("adb")
	scrcpyPath := a.getToolPath("scrcpy")
	if _, err := os.Stat(scrcpyPath); os.IsNotExist(err) {
		return "Error: scrcpy not found at " + scrcpyPath
	}

	err := mirror.StartMirror(adbPath, scrcpyPath, deviceID, "Mirror: "+deviceID)
	if err != nil {
		return "Error: " + err.Error()
	}
	return "Started"
}
