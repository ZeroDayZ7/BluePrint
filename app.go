package main

import (
	"BluePrint/backend/adb"
	"BluePrint/backend/mirror"
	"context"
	"log"
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
	exePath, err := os.Executable()
	if err != nil {
		log.Println("ERROR: cannot get executable path:", err)
		return ""
	}

	exeDir := filepath.Dir(exePath)
	log.Println("Executable dir:", exeDir)

	if toolName == "adb" || toolName == "scrcpy" {
		toolName += ".exe"
	}

	fullPath := filepath.Join(exeDir, "adb_tools", toolName)

	log.Println("Resolved tool path:", fullPath)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		log.Println("ERROR: tool does not exist:", fullPath)
	}

	return fullPath
}

func (a *App) GetDevices() []Device {
	adbPath := a.getToolPath("adb")
	log.Println("DEBUG: Próba pobrania urządzeń używając:", adbPath)

	rawOutput, err := adb.FetchDevices(adbPath)

	if err != nil {
		log.Println("ERROR: Błąd FetchDevices:", err)
		return []Device{}
	}

	log.Println("DEBUG: Raw ADB Output:", rawOutput)

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
	log.Printf("DEBUG: Przetworzono urządzeń: %d", len(devices))
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

func (a *App) GetInstalledApps(deviceID string, onlyUserApps bool) []string {
	adbPath := a.getToolPath("adb")

	filter := "-s"
	if onlyUserApps {
		filter = "-3"
	}

	rawOutput, err := adb.FetchPackages(adbPath, deviceID, filter)
	if err != nil {
		log.Println("ERROR:", err)
		return []string{}
	}

	var packages []string
	lines := strings.Split(rawOutput, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			packages = append(packages, strings.TrimPrefix(line, "package:"))
		}
	}
	return packages
}
