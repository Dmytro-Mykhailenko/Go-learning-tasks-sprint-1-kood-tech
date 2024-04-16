package sprint

func ToUpperCase(s string) string {
	var str string
	for _, ch := range s {
		if IsLower(string(ch)) {
			str += string(ch - 32)
			continue
		}
		str += string(ch)
	}
	return str
}
