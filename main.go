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

	classRoom2 := classRoom
	delete(classRoom2, "1/1")

	fmt.Println(classRoom, classRoom2)
}
