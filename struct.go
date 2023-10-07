package main

import "fmt"

type car struct {
	name string
	year int
}

func newCar(car car) *car {
	car.year = 32
	fmt.Print(car)
	return &car

}

func main() {
	newCar(car{name: "tp", year: 50})
	newCar(car{"toyota", 24})
	fmt.Print(car{"tpyp", 23})
}
