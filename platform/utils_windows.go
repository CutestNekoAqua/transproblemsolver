//go:build windows
// +build windows

package platform

import "github.com/kindlyfire/go-keylogger"

const PATH_SEPARATOR = '\\'

var logger = keylogger.NewKeylogger()

func GetKey() (bool, string) {
	key := logger.GetKey()
	if key.Empty {
		return true, ""
	}
	return false, string(key.Rune)
}
