package sprint

func Abacus(a int, b int) int {

	if b == 0 {
		panic("Error. Division by zero")
	}

	c := a / b

	return c

}
