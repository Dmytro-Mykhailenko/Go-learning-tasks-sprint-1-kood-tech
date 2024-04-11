package sprint

func AlphabetMastery(n int) string {

	ch := 'a'

	var resStr string

	for ; n > 0; n-- {

		resStr += string(ch)
		ch += 1

	}

	return resStr

}
