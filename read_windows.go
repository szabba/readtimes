package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

//go:embed tts.ps1
var psScript string

func read(text string) error {
	cmd := exec.Command("powershell", "-Command", psScript)

	cmd.Env = []string{
		fmt.Sprintf("TTS_TEXT=%q", text),
	}

	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
