package mirror

import (
	"fmt"
	"os/exec"
)

// StartMirror uruchamia scrcpy dla konkretnego urządzenia
func StartMirror(adbPath string, scrcpyPath string, deviceID string, title string) error {
	// Flagi:
	// -s: konkretne urządzenie
	// --window-title: własny tytuł okna
	// --always-on-top: opcjonalnie
	cmd := exec.Command(scrcpyPath,
		"-s", deviceID,
		"--window-title", title,
		"--always-on-top",
	)

	// Uruchamiamy w tle, nie czekamy na zakończenie (Start zamiast Run)
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start scrcpy: %v", err)
	}

	return nil
}
