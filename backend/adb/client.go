// backend/adb/client.go
package adb

import "os/exec"

func FetchDevices(adbPath string) (string, error) {
	cmd := exec.Command(adbPath, "devices", "-l")
	out, err := cmd.Output()
	return string(out), err
}

func FetchPackages(adbPath string, deviceID string, filter string) (string, error) {
	args := []string{"-s", deviceID, "shell", "pm", "list", "packages"}

	if filter != "" {
		args = append(args, filter)
	}

	cmd := exec.Command(adbPath, args...)
	out, err := cmd.Output()
	return string(out), err
}
