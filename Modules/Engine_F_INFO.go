package GRIP

import (
	"fmt"
	"os"
)

func FINF(filename string) {
	f, x := os.Stat(filename)
	if x != nil {
		fmt.Println("(ERROR FATAL) ======================== ", x)
	}
	fmt.Println("\t FILE NAME       \n\t| ", f.Name())
	fmt.Println("\t FILE MOD TIME   \n\t| ", f.ModTime())
	fmt.Println("\t FILE MODE       \n\t| ", f.Mode())
	fmt.Println("\t FILE IS DIR     \n\t| ", f.IsDir())
	fmt.Println("\t FILE SIZE       \n\t| ", f.Size())
}
