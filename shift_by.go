package sprint

func ShiftBy(r rune, step int) rune {

	var b rune

	if r+rune(step) > 'z' {

		b = (r - 26) + rune(step)

	} else {
		b = r + rune(step)
	}

	return b

}
