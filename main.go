package main

import (
	"fmt"
)

func main() {
	classRoom := make(map[string]int)
	classRoom = map[string]int{
		"1/1": 20,
		"1/2": 22,
		"1/3": 23,
	}

	fmt.Printf("%v \n", classRoom)
}
