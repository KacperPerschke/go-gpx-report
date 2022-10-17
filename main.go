package main

import (
	"fmt"
)

func main() {
	conf, err := prodConf()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", conf)
	if conf.ReadFromFile() != "" {
		fmt.Println("From file")
	}
	if conf.ReadFromStdin() {
		fmt.Println("From STDIN")
	}
	if err != nil {
		panic(err)
	}
}
