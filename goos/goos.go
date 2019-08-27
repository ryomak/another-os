package main

import (
	"fmt"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("this computer is windows")
	} else {
    fmt.Println("this computer is another :",runtime.GOOS)
	}
}
