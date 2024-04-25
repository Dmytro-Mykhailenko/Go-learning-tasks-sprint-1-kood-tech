package sprint

func DigitalRoot(n int) int {

	sum := 0

	for n > 0 {

		y := n % 10
		sum += y
		n /= 10

		if n == 0 && sum > 10 {
			n = sum
			sum = 0
		}
	}

	return sum
}
