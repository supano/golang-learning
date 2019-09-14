package main

import (
	"fmt"
	"strconv"
)

var (
	i int     = 100
	j float64 = 200.50
)

func main() {
	fmt.Printf("%v, %T \r\n", i, i)

	k := strconv.Itoa(i)
	fmt.Printf("%v, %T", k, k)
}
