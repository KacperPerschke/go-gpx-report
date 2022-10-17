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
}
