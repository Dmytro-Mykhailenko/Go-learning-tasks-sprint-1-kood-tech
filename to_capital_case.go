package sprint

func ToCapitalCase(s string) string {
	res := []rune(s)
	var firstSimbolOrDigit bool

	for i, ch := range res {

		if IsSimbolOrDigit(ch) && !firstSimbolOrDigit {
			firstSimbolOrDigit = true
			if IsSimbol(ch) {
				if IsLower(ch) {
					res[i] = ch - 32
				}
			}
		}

		if !IsSimbolOrDigit(ch) {
			firstSimbolOrDigit = false
		}

	}
	return string(res)
}

func IsLower(ch rune) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	return false
}

func IsSimbol(ch rune) bool {
	if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
		return true
	}
	return false
}

func IsDigit(ch rune) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}

func IsSimbolOrDigit(ch rune) bool {
	if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' {
		return true
	}
	return false
}
