package GRIP

import (
	"bufio"
	"fmt"
	"os"
)

type Banner struct {
	FNAME string // File name
	FVERS string // Version
}

func (Text *Banner) Output_banner() {
	fmt.Println("\x1b[H\x1b[2J\x1b[3J")
	f, x := os.Open(Text.FNAME)
	TextErrors.Exit = false
	TextErrors.Y = " \033[31mCould not open the banner file, it is suggested to fix this, file should be located in Modules/Text but it isnt or something went wrong - \033[39m"
	TextErrors.X = x
	fmt.Println("\033[39m")
	TextErrors.ExitC = 0
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
