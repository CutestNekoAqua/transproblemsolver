//go:build !windows
// +build !windows

package platform

import (
	"fmt"
	"github.com/MarinX/keylogger"
	"os"
)

const PATH_SEPARATOR = '/'

var logger keylogger.KeyLogger

func init() {
	fmt.Print("Searching Keyboard..")
	keyboard := keylogger.FindKeyboardDevice()
	if keyboard == "" {
		fmt.Println("No keyboard found! Exiting")
		os.Exit(1)
	}
	fmt.Println(keyboard)
	logger = keylogger.New(keyboard)
}

func GetKey() (bool, string) {
read:
	keylogger.InputEvent = logger.Read()
key:
	string = read.KeyString()
	if key == "" {
		return true, ""
	}
	return false, key
}
