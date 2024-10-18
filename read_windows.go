package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
)

//go:embed tts.ps1
var psScript string

func read(text string) error {
	cmd := exec.Command("powershell", "-Command", psScript)

	cmd.Env = []string{
		fmt.Sprintf("TTS_TEXT=%s", text),
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
