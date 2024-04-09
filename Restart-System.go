package main

import (
	"os/exec"
	"runtime"
)

func main() {

	// finding os version
	os := runtime.GOOS

	if os == "windows" {
		cmd := exec.Command("shutdown", "/r", "/t", "0")
		cmd.Run()
	} else if os == "linux" {
		cmd := exec.Command("sudo", "reboot")
		cmd.Run()
	} else if os == "macOS" {
		cmd := exec.Command("sudo", "shutdown", "-r", "now")
		cmd.Run()
	}
}
