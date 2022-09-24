package GRIP

import (
	"fmt"
	"os"
	"time"
)

type Error struct {
	X     error     // Error
	Y     string    // Message to output
	D     time.Time // Time to output
	Exit  bool      // Exit or
	ExitC int       // if exit what code?
}

func (E *Error) CE() {
	if E.X != nil {
		fmt.Println("[ERROR] :::>> ", E.Y, E.X)
		if E.Exit {
			os.Exit(E.ExitC)
		}
	}
}
