package sprint

func FactorialRecursive(n int) int {

	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else if n == 1 || n == 2 {
		return n
	}

	x := FactorialRecursive(n - 1)

	if y := x * n; y <= 0 {
		return 0
	} else {
		return y
	}
}
