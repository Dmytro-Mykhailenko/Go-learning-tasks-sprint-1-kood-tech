package sprint

func ShiftBy(r rune, step int) rune {

	var b rune

	b = r >> step

	return b

}
