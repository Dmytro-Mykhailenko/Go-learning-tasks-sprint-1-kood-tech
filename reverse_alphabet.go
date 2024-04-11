package sprint

func ReverseAlphabet(step int) string {

	if step <= 0 {

		step = 1

	}

	ch := 'z'

	var resStr string

	for ch >= 'a' {

		resStr += string(ch)

		ch -= rune(step)

	}

	return resStr

}
