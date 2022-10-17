package main

import (
	"fmt"
	"os"
)

func main() {
	file := os.Stdin
	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat()", err)
	}
	size := fi.Size()
	if size > 0 {
		fmt.Printf("%v bytes available in Stdin\n", size)
	} else {
		fmt.Println("Stdin is empty")
	}
}
