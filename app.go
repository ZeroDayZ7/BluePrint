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

type ProcessInfo struct {
	PID  string `json:"pid"`
	User string `json:"user"`
	CPU  string `json:"cpu"`
	Mem  string `json:"mem"`
	Name string `json:"name"`
}

type FileEntry struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
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

func (a *App) GetProcesses(deviceID string) []ProcessInfo {
	adbPath := a.getToolPath("adb")
	log.Printf("DEBUG: Start GetProcesses for device: %s", deviceID)

	raw, err := adb.FetchProcesses(adbPath, deviceID)
	if err != nil {
		log.Println("ERROR GetProcesses (ADB Fetch):", err)
		return []ProcessInfo{}
	}

	// LOG 1: Zobaczmy ile znaków i linii dostaliśmy w ogóle
	log.Printf("DEBUG: Raw output length: %d chars", len(raw))

	var processes []ProcessInfo
	lines := strings.Split(raw, "\n")

	startParsing := false
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// LOG 2: Logujemy pierwsze 10 linii, żeby zobaczyć jak wygląda nagłówek topa
		if i < 15 {
			log.Printf("DEBUG: Line %d: [%s]", i, line)
		}

		// Próbujemy wykryć nagłówek bardziej elastycznie
		// Niektóre wersje Androida mają PID, USER, COMMAND, inne PID, USER, NAME, inne jeszcze ARGS
		if !startParsing && strings.Contains(line, "PID") && (strings.Contains(line, "NAME") || strings.Contains(line, "ARGS") || strings.Contains(line, "CMD") || strings.Contains(line, "COMMAND")) {
			log.Printf("DEBUG: Found header at line %d: %s", i, line)
			startParsing = true
			continue
		}

		if startParsing {
			fields := strings.Fields(line)
			if len(fields) >= 9 {
				name := fields[len(fields)-1]

				processes = append(processes, ProcessInfo{
					PID:  fields[0],
					User: fields[1],
					CPU:  fields[8],
					Mem:  fields[9],
					Name: name,
				})
			}
		}
	}

	log.Printf("DEBUG: Finished parsing. Found %d processes.", len(processes))

	// LOG 3: Jeśli nic nie znaleźliśmy, a mieliśmy dane, to znaczy że startParsing zawiodło
	if len(processes) == 0 && len(lines) > 5 {
		log.Println("WARNING: Data was received but no processes were parsed. Check if 'startParsing' became true.")
	}

	return processes
}

func (a *App) KillProcess(deviceID string, pid string) string {
	log.Printf("DEBUG: Attempting to kill PID %s on device %s", pid, deviceID)
	adbPath := a.getToolPath("adb")
	err := adb.KillProcess(adbPath, deviceID, pid)
	if err != nil {
		log.Printf("ERROR KillProcess: %v", err)
		return "Error: " + err.Error()
	}
	log.Printf("DEBUG: Kill signal sent successfully to PID %s", pid)
	return "Success"
}

func (a *App) ListFiles(deviceID string, path string) []FileEntry {
	adbPath := a.getToolPath("adb")

	// Pobieramy surowe dane z modułu adb
	raw, err := adb.FetchFiles(adbPath, deviceID, path)
	if err != nil {
		log.Println("ERROR ListFiles:", err)
		return []FileEntry{}
	}

	var files []FileEntry
	lines := strings.Split(raw, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Pomijamy puste linie i błędy uprawnień
		if line == "" || strings.Contains(line, "Permission denied") {
			continue
		}

		// Jeśli nazwa kończy się na /, to jest to katalog (dzięki flagi -p w ls)
		isDir := strings.HasSuffix(line, "/")
		name := strings.TrimSuffix(line, "/")

		files = append(files, FileEntry{
			Name:  name,
			IsDir: isDir,
		})
	}

	log.Printf("DEBUG: Found %d files in %s", len(files), path)
	return files
}
