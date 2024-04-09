package sprint

func ReverseAlphabetValue(ch rune) rune {

	var b rune

	if 'z'-ch > 13 {

		b = 'z' - (ch - 'a')

	} else {
		b = 'a' + ('z' - ch)
	}

	return b

}
