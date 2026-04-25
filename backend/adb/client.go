package adb

import "os/exec"

func FetchDevices(adbPath string) (string, error) {
	cmd := exec.Command(adbPath, "devices", "-l")
	out, err := cmd.Output()
	return string(out), err
}
