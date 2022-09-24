package GRIP

import (
	"fmt"
	"os"
)

func (F *File) Delete() {
	x := os.Remove(F.Filename)
	TextErrors.Exit = false
	TextErrors.ExitC = 0
	TextErrors.X = x
	TextErrors.Y = " Could not remove the filename, got an error of prevention -> "
	TextErrors.CE()
	if x == nil {
		fmt.Println(RED, "[INFO] :::: File deleted - ", F.Filename)
	}
}
