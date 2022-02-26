package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
	"log"
	"os"
	"strings"
	"time"
	"transProblemsSolver/platform"
)

const (
	delayKeyfetchMS = 5
)

func main() {
	a, err := os.ReadFile("replacer.txt")
	lines := strings.Split(string(a), "\r\n")
	var confs = make([][2]string, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ":")
		confs[i][0] = parts[0]
		confs[i][1] = parts[1]
	}
	if err != nil {
		log.Fatal(err)
	}
	current := ""
	for {
		empty, key := platform.GetKey()
		if !empty {
			if key == " " {
				current = ""
				continue
			}
			i := loopOptions(confs, current, key)
			if i != -1 {
				current += key
				shortcut := confs[i][0]
				if shortcut == current {
					current = ""
					fmt.Println("found", shortcut)
					replace(confs[i][1], len(shortcut))
				}
			} else {
				current = ""
			}
		}
		time.Sleep(delayKeyfetchMS * time.Millisecond)
	}

}

func loopOptions(array [][2]string, current string, key string) int {
	currentBiggerAll := true
	for i, line := range array {
		shortcut := line[0]
		if len(shortcut) >= len(current+key) {
			currentBiggerAll = false
		}
		if strings.Contains(shortcut, current+key) && !currentBiggerAll {
			return i
		}
	}
	return -1
}

func replace(with string, length int) {
	for i := 0; i < length; i++ {
		time.Sleep(delayKeyfetchMS * time.Millisecond * 2)
		robotgo.KeyTap("left")
		time.Sleep(delayKeyfetchMS * time.Millisecond * 2)
		robotgo.KeyTap("delete")
	}
	content, _ := clipboard.ReadAll()
	robotgo.PasteStr(with)
	time.Sleep(delayKeyfetchMS * time.Millisecond)
	err := clipboard.WriteAll(content)
	if err != nil {
		log.Fatal(err)
	}
}
