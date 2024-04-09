package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num1 := os.Args[1]
	num2 := os.Args[2]

	a, _ := strconv.Atoi(num1)
	b, _ := strconv.Atoi(num2)

	fmt.Println(Abacus(a, b))

}

func Abacus(a int, b int) int {

	if b == 0 {
		panic("Error. Division by zero")
	}

	c := a / b

	return c

}
