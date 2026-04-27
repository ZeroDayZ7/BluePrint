package mirror

import (
	"fmt"
	"os/exec"
	"syscall"
)

func StartMirror(adbPath string, scrcpyPath string, deviceID string, title string) error {
	cmd := exec.Command(scrcpyPath,
		"-s", deviceID,
		"--window-title", title,
		"--always-on-top",
	)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000,
	}

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start scrcpy: %v", err)
	}

	return nil
}
