package sprint

func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {
	number := ConvertAnyToDec(nbr, baseFrom)
	return NbrBase(number, baseTo)
}
func ConvertAnyToDec(s string, base string) int {

	if !IsValidBase(base) {
		return 0
	}

	x := len(base)
	result := 0

	for i, char := range s {
		index := indexOf(char, base)
		if index == -1 {
			return 0
		}
		y := len(s) - i - 1
		result += index * powerOf(x, y)
	}

	return result
}

func powerOf(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	return powerOf(x, n-1) * x
}

func indexOf(char rune, base string) int {
	for i, b := range base {
		if b == char {
			return i
		}
	}
	return -1
}

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
