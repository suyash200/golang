package main

import "fmt"

func main() {

	v := make(map[string]int, 0)
	v["k1"] = 234
	v["k1"] = 7
	v["k2"] = 13

	fmt.Println("map:", v)
}
