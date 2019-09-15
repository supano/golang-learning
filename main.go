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

	fmt.Println(student1.favFood, student2.favFood)
}
