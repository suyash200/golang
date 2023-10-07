package main

import (
	"fmt"
)

type rect struct {
	length  int8
	breadth int8
}

func (rect rect) area() int16 {
	return int16(rect.length) * int16(rect.breadth)
}

func main() {
	r := rect{12, 12}
	fmt.Println(r.area())

	rp := &r
	fmt.Println(rp.area())

}
