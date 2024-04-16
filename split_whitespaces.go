package sprint

func SplitWhitespaces(s string) []string {
	var res []string
	var word []rune

	for i, ch := range s {

		if ch != 9 && ch != 10 && ch != 32 || i == len(s)-1 {
			word = append(word, ch)
		}

		if ch == 9 || ch == 10 || ch == 32 || i == len(s)-1 {
			res = append(res, string(word))
			word = []rune{}
		}
	}
	return res
}
