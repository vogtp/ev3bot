package gobot

import (
	"fmt"
	"os/exec"
)

type display struct {
}

func createDisplay() *display {
	d := &display{}
	//d.SetFont("Lat15-Fixed14")
	d.SetFont("Lat15-Fixed16")
	d.ShowCursor(false)
	return d
}

func (d *display) ShowCursor(on bool) {
	// use escape code to turn cursor on or off
	if on {
		fmt.Print("\x1B[?25h")
	} else {
		fmt.Print("\x1B[?25l")
	}
}

func (d *display) SetFont(name string) {
	// use setfont command to change the font
	// run `ls /usr/share/consolefonts` in a terminal to get a list of available fonts
	exec.Command("setfont", name).Run()
}

func (d *display) Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

func (d *display) Print(a ...interface{}) (n int, err error) {
	return fmt.Print(a...)
}
