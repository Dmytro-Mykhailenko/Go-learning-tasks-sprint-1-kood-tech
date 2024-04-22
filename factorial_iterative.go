package sprint

func FactorialIterative(n int) int {

	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else if n == 1 || n == 2 {
		return n
	}

	var x = 2

	for i := 3; i <= n; i++ {
		x *= i
		if x <= 0 {
			return 0
		}
	}

	return x
}
