package keyboard

import (
	"fmt"
	"os"
)

const ledPath = "/sys/class/leds/dell::kbd_backlight/brightness"

// const ledPath = "./brightness"

// Toggle toggles the keyboard backlight
func Toggle() {
	file, err := os.Open(ledPath)
	// file, err := os.Open(ledPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	buff := make([]byte, 1)
	_, err = file.Read(buff)
	if err != nil {
		fmt.Println(err)
		return
	}
	backLightVal := string(buff)
	fmt.Println(backLightVal)
	if backLightVal != "0" {
		backLightVal = "0"
	} else {
		backLightVal = "1"
	}
	// file, err = os.OpenFile(ledPath, os.O_RDONLY|os.O_WRONLY|os.O_TRUNC, 0777)
	file, err = os.OpenFile(ledPath, os.O_RDONLY|os.O_WRONLY|os.O_TRUNC, os.ModeAppend)

	if err != nil {
		fmt.Println("OPEN ERR", err)
		return
	}
	_, err = file.Write([]byte(backLightVal))
	if err != nil {
		fmt.Println("WRITE ERR", err)
		return
	}
	fmt.Println([]byte(backLightVal))
}
