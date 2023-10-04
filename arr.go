package main

import (
	"fmt"
)

func main() {

	var arrlength int
	var arr1 [5]int
	fmt.Scan(&arrlength)

	for i := 0; i <= arrlength; i++ {
		var inp int
		fmt.Scan(&inp)
		arr1[i] = inp
	}

	for i := 0; i <= arrlength; i++ {
		fmt.Println("element is ", arr1[i])
	}

}
