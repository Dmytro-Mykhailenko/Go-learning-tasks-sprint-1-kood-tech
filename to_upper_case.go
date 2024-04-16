package sprint

func ToUpperCase(s string) string {
	var str string
	for _, ch := range s {
		if IsLower(ch) {
			str += string(ch - 32)
			continue
		}
		str += string(ch)
	}
	return str
}

func IsLower(ch rune) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	return false
}
