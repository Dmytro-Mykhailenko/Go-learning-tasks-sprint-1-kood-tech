package sprint

func ToCapitalCase(s string) string {
	res := []rune(s)
	var firstSimbolOrDigit bool

	for i, ch := range res {

		if IsSimbolOrDigit(ch) && !firstSimbolOrDigit {
			firstSimbolOrDigit = true
			if IsSimbol(ch) && IsLower(ch) {
				res[i] = LowerToUpper(ch)
			}
		} else if IsSimbol(ch) && firstSimbolOrDigit {
			if IsLower(ch) {
				res[i] = LowerToUpper(ch)
			} else if IsUpper(ch) {
				res[i] = UpperToLower(ch)
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

func IsUpper(ch rune) bool {
	if ch >= 'A' && ch <= 'Z' {
		return true
	}
	return false
}

func LowerToUpper(ch rune) rune {
	return ch - 32
}

func UpperToLower(ch rune) rune {
	return ch + 32
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
