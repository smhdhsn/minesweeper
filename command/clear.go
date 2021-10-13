package command

import (
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func()

// init, initializes the functionality of the clear screen command for [Linux, Windows] operating systems.
func init() {
	clear = make(map[string]func())

	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// Clear, clears the screen.
func Clear() {
	value, ok := clear[runtime.GOOS]

	if ok {
		value()
	} else {
		panic("undefined platform!")
	}
}
