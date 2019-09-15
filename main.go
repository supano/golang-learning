package main

import "fmt"

type student struct {
	name    string
	id      int
	favFood []string
}

func main() {
	student1 := student{
		name: "student1",
		id:   1,
		favFood: []string{
			"burger",
			"pizza",
		},
	}

	student2 := student{
		name: "student2",
		id:   2,
		favFood: []string{
			"sushi",
			"okonomiyagi",
		},
	}

	if student1.name == "student1" && student2.name == "student2" {
		fmt.Println(student1.favFood, student2.favFood)
	}

	switch student1.favFood[0] {
	case "burger":
		fmt.Println("he's like burger")
	case "pizza":
		fmt.Println("he's like pizza")
	default:
		fmt.Println("he's like nothing")
	}
}
