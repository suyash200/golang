package main

import (
	"fmt"
)

func main() {
	fmt.Print("hellow")
	var v make([]string,0)
	fmt.Println("uninit:", v, v == nil, len(v) == 0)

	for i := 0; i <= 6; i++ {
		fmt.Print(v[i])
	}
}
