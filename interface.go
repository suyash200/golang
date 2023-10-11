package main

import (
	"fmt"
	"math"
)

type rect interface {
	area() float64
	permiter() float64
}

func main() {
	fmt.Println(math.Pi)
}
