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
	shortcut        = "<3"
)

func main() {
	emoji, err := os.ReadFile("replacer.txt")
	if err != nil {
		log.Fatal(err)
	}
	current := ""
	for {
		empty, key := platform.GetKey()
		if !empty {
			if strings.Contains(shortcut, current+key) {
				current += key
				if shortcut == current {
					current = ""
					fmt.Println("found", shortcut)
					replace(string(emoji), len(shortcut))
				}
			} else {
				current = ""
			}
		}
		time.Sleep(delayKeyfetchMS * time.Millisecond)
	}

}

func replace(with string, length int) {
	for i := 0; i < length; i++ {
		robotgo.KeyTap("left")
		time.Sleep(delayKeyfetchMS * time.Millisecond * 5)
		robotgo.KeyTap("delete")
		time.Sleep(delayKeyfetchMS * time.Millisecond * 5)
	}
	content, _ := clipboard.ReadAll()
	robotgo.PasteStr(with)
	time.Sleep(delayKeyfetchMS * time.Millisecond)
	clipboard.WriteAll(content)
}
