package sprint

func LCM(a, b int) int {

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

	x, y := a, b

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

	c := (x * y) / a

	return c
}
