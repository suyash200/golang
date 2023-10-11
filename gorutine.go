package main

import (
	"fmt"
	"time"
)

func m(he string) {
	fmt.Print(he)
}

func main() {
	m("jhello")
	go m("hellow")
	fmt.Print("hellow")
	time.Sleep(time.Second)
	fmt.Println("done")
}
