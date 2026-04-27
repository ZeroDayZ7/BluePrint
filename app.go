package main

import (
	"BluePrint/backend/adb"
	"BluePrint/backend/config"
	"BluePrint/backend/mirror"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx    context.Context
	Config *config.AppConfig
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

type DeviceProp struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DeviceInfoResponse struct {
	Model        string       `json:"model"`
	Manufacturer string       `json:"manufacturer"`
	AndroidVer   string       `json:"androidVer"`
	APILevel     string       `json:"apiLevel"`
	CPU          string       `json:"cpu"`
	Serial       string       `json:"serial"`
	Battery      string       `json:"battery"`
	AllProps     []DeviceProp `json:"allProps"`
}

func NewApp() *App {
	return &App{
		Config: config.Load(),
	}
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

func (a *App) ListFiles(deviceID string, path string, showHidden bool) []FileEntry {
	adbPath := a.getToolPath("adb")

	raw, err := adb.FetchFiles(adbPath, deviceID, path, showHidden)

	log.Printf("DEBUG ADB RAW for %s: [%s]", path, raw)

	// ROZWIĄZANIE: Przerywamy tylko wtedy, gdy jest błąd ORAZ nie ma żadnych danych w `raw`
	if err != nil && strings.TrimSpace(raw) == "" {
		log.Println("ERROR ListFiles (no data):", err)
		return []FileEntry{}
	}
	// Jeśli err != nil, ale raw ma dane, ignorujemy błąd i lecimy dalej!

	var files []FileEntry
	lines := strings.Split(raw, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.Contains(line, "Permission denied") || strings.Contains(line, "No such file") {
			continue
		}

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

func (a *App) DownloadFile(deviceID string, remotePath string, fileName string) string {
	adbPath := a.getToolPath("adb")

	localPath, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultFilename: fileName,
		Title:           "Gdzie zapisać plik?",
	})

	if err != nil || localPath == "" {
		return "Cancelled"
	}

	err = adb.PullFile(adbPath, deviceID, remotePath, localPath)
	if err != nil {
		return "Error: " + err.Error()
	}
	return "Success"
}

func (a *App) DeleteFile(deviceID string, path string) bool {
	adbPath := a.getToolPath("adb")
	err := adb.RemoveItem(adbPath, deviceID, path)
	if err != nil {
		log.Printf("ERROR: Delete failed: %v", err)
		return false
	}
	return true
}

func (a *App) GetStoragePoints(deviceID string) []string {
	adbPath := a.getToolPath("adb")
	raw, err := adb.FetchStoragePoints(adbPath, deviceID)

	storageList := []string{"/sdcard"}

	if err != nil {
		log.Println("ERROR GetStoragePoints:", err)
		return storageList
	}

	lines := strings.Split(raw, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// ROZWIĄZANIE: Dodano warunek strings.Contains(line, "enc_")
		if line == "" || line == "self" || line == "emulated" || line == "container" || strings.Contains(line, "enc_") {
			continue
		}

		storageList = append(storageList, "/storage/"+line)
	}

	log.Printf("DEBUG: Found storage points for %s: %v", deviceID, storageList)
	return storageList
}

func (a *App) GetDeviceInfo(deviceID string) DeviceInfoResponse {
	fmt.Println("!!! FRONTEND MNIE WYWOŁAŁ !!! ID:", deviceID)
	log.Printf("INFO: Starting GetDeviceInfo for ID: %s", deviceID)
	adbPath := a.getToolPath("adb")
	data := DeviceInfoResponse{
		Battery: "Unknown",
	}

	rawProps, err := adb.FetchDeviceProps(adbPath, deviceID)
	if err != nil {
		log.Printf("ERROR: GetDeviceInfo/FetchDeviceProps: %v", err)
	} else {
		lines := strings.Split(rawProps, "\n")
		propMap := make(map[string]string)

		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			parts := strings.SplitN(line, ": ", 2)
			if len(parts) == 2 {
				key := strings.Trim(parts[0], "[]")
				val := strings.Trim(parts[1], "[]")
				propMap[key] = val
				data.AllProps = append(data.AllProps, DeviceProp{Key: key, Value: val})
			}
		}

		data.Model = propMap["ro.product.model"]
		data.Manufacturer = propMap["ro.product.manufacturer"]
		data.AndroidVer = propMap["ro.build.version.release"]
		data.APILevel = propMap["ro.build.version.sdk"]
		data.CPU = propMap["ro.product.cpu.abi"]
		data.Serial = propMap["ro.serialno"]

		log.Printf("DEBUG: Parsed %d properties for model: %s", len(data.AllProps), data.Model)
	}

	rawBat, err := adb.FetchBatteryStatus(adbPath, deviceID)
	if err != nil {
		log.Printf("ERROR: GetDeviceInfo/FetchBatteryStatus: %v", err)
	} else {
		lines := strings.Split(rawBat, "\n")
		foundBattery := false
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "level:") {
				parts := strings.Split(line, ":")
				if len(parts) == 2 {
					data.Battery = strings.TrimSpace(parts[1]) + "%"
					log.Printf("DEBUG: Battery level found: %s", data.Battery)
					foundBattery = true
				}
				break
			}
		}
		if !foundBattery {
			log.Printf("WARN: Battery level not found in dumpsys output")
		}
	}

	log.Printf("INFO: Finished GetDeviceInfo for %s (%s)", data.Model, deviceID)
	return data
}

func (a *App) ExecuteShell(deviceID string, command string) string {
	if deviceID == "" {
		return "Error: No device selected"
	}
	adbPath := a.getToolPath("adb")
	output, err := adb.RunShellCommand(adbPath, deviceID, command)
	if err != nil && output == "" {
		return "Error: " + err.Error()
	}
	return output
}

func (a *App) GetSettings() *config.AppConfig {
	return a.Config
}

func (a *App) SaveSettings(newConfig config.AppConfig) string {
	*a.Config = newConfig

	err := config.Save(a.Config)
	if err != nil {
		return err.Error()
	}
	return "Success"
}
