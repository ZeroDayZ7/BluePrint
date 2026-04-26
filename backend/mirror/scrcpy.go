package mirror

import (
	"fmt"
	"os/exec"
)

func StartMirror(adbPath string, scrcpyPath string, deviceID string, title string) error {
	cmd := exec.Command(scrcpyPath,
		"-s", deviceID,
		"--window-title", title,
		"--always-on-top",
	)

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start scrcpy: %v", err)
	}

	return nil
}
