package main

import "fmt"

func trial[k comparable, v int8 | float64](m map[k]v) {
	fmt.Print(m)
}
