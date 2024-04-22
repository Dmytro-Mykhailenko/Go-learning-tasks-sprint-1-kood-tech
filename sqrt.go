package sprint

func Sqrt(n int) int {

	if n == 0 || n == 1 {
		return n
	}

	z := n
	var x int

	for z >= 0 {
		x = z / 2
		y := ToThePowerIterative(x, 2)
		if y == n {
			return x
		} else if y > n {
			z--
		} else if y < n {
			return 0
		}
	}
	return 0
}
