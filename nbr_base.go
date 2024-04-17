package sprint

func NbrBase(n int, base string) string {

	if !IsValidBase(base) {
		return "NV"
	}

	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	x := len(base)
	var result string

	for n > 0 {
		y := n % x
		result = string(base[y]) + result
		n /= x
	}

	if isNegative {
		result = "-" + result
	}

	return result
}

func IsValidBase(s string) bool {

	if len(s) < 2 {
		return false
	}

	var uniqueChars string
	for _, char := range s {
		found := false
		if char == '-' || char == '+' {
			return false
		}
		for _, existingChar := range uniqueChars {

			if existingChar == char {
				found = true
				break
			}
		}
		if !found {
			uniqueChars += string(char)
		}
	}
	if len(s) != len(uniqueChars) {
		return false
	}
	return true
}
