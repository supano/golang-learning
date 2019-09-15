package main

import (
	"fmt"
)

func main() {
	grades := []int{1, 2, 3}
	a := grades[:]
	b := grades[1:]

	fmt.Printf("a: %v \n", a)
	fmt.Printf("b: %v \n", b)
}
