package adb

import (
	"log"
	"os/exec"
	"syscall"
)

func silentCmd(cmd *exec.Cmd) *exec.Cmd {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000,
	}
	return cmd
}

func FetchDevices(adbPath string) (string, error) {
	cmd := silentCmd(exec.Command(adbPath, "devices", "-l"))
	out, err := cmd.Output()
	return string(out), err
}

func FetchPackages(adbPath string, deviceID string, filter string) (string, error) {
	args := []string{"-s", deviceID, "shell", "pm", "list", "packages"}
	if filter != "" {
		args = append(args, filter)
	}
	cmd := silentCmd(exec.Command(adbPath, args...))
	out, err := cmd.Output()
	return string(out), err
}

func FetchProcesses(adbPath string, deviceID string) (string, error) {
	cmd := silentCmd(exec.Command(adbPath, "-s", deviceID, "shell", "top", "-b", "-n", "1"))
	out, err := cmd.Output()
	return string(out), err
}

func KillProcess(adbPath string, deviceID string, pid string) error {
	cmd := silentCmd(exec.Command(adbPath, "-s", deviceID, "shell", "kill", "-9", pid))
	return cmd.Run()
}

func FetchFiles(adbPath string, deviceID string, path string, showHidden bool) (string, error) {
	lsCmd := "ls -p"
	if showHidden {
		lsCmd = "ls -ap"
	}
	args := []string{"-s", deviceID, "shell", lsCmd, path}
	cmd := silentCmd(exec.Command(adbPath, args...))
	out, err := cmd.Output()
	return string(out), err
}

func PullFile(adbPath, deviceID, remotePath, localPath string) error {
	cmd := silentCmd(exec.Command(adbPath, "-s", deviceID, "pull", remotePath, localPath))
	return cmd.Run()
}

func RemoveItem(adbPath, deviceID, path string) error {
	cmd := silentCmd(exec.Command(adbPath, "-s", deviceID, "shell", "rm", "-rf", path))
	return cmd.Run()
}

func FetchStoragePoints(adbPath string, deviceID string) (string, error) {
	cmd := silentCmd(exec.Command(adbPath, "-s", deviceID, "shell", "ls", "/storage"))
	out, err := cmd.Output()
	return string(out), err
}

func FetchDeviceProps(adbPath string, deviceID string) (string, error) {
	log.Printf("DEBUG: Executing getprop for device: %s", deviceID)
	cmd := silentCmd(exec.Command(adbPath, "-s", deviceID, "shell", "getprop"))
	out, err := cmd.Output()
	if err != nil {
		log.Printf("ERROR: FetchDeviceProps failed: %v", err)
	}
	return string(out), err
}

func FetchBatteryStatus(adbPath string, deviceID string) (string, error) {
	log.Printf("DEBUG: Executing dumpsys battery for device: %s", deviceID)
	cmd := silentCmd(exec.Command(adbPath, "-s", deviceID, "shell", "dumpsys", "battery"))
	out, err := cmd.Output()
	if err != nil {
		log.Printf("ERROR: FetchBatteryStatus failed: %v", err)
	}
	return string(out), err
}

func RunShellCommand(adbPath string, deviceID string, command string) (string, error) {
	cmd := silentCmd(exec.Command(adbPath, "-s", deviceID, "shell", command))
	out, err := cmd.CombinedOutput()
	return string(out), err
}
