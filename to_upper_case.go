package sprint

import "unicode"

func ToUpperCase(s string) string {
	var str string
	for _, ch := range s {
		if unicode.IsLetter(ch) {
			if unicode.IsLower(ch) {
				str += string(ch - 32)
				continue
			}
		}
		str += string(ch)
	}
	return str
}
