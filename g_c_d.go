package sprint

func GCD(a, b int) int {

	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	for a != b {
		if a <= 0 || b <= 0 {
			return 1
		}
		if a < b {
			b -= a
		}
		if a > b {
			a -= b
		}
	}

	return a
}
