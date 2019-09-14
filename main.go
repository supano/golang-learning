package main

import (
	"fmt"
)

var (
	i int     = 100
	j float64 = 200.50
)

func main() {
	const myConst string = "my first constant"
	fmt.Printf("%v, %T", myConst, myConst)
}
