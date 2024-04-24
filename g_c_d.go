package sprint

func GCD(a, b int) int {

	if a < 0 {
		a = -a
	}

	if b < 0 {
		b = -b
	}

	if a == 0 {
		return b
	}

	if b == 0 {
		return a
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
