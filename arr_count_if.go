package sprint

func ArrCountIf(f func(string) bool, tab []string) int {
	var cnt int

	for _, str := range tab {
		if f(str) {
			cnt++
		}
	}
	return cnt
}

func IsUpper(s string) bool {
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			continue
		}
		return false
	}
	return true
}

func IsAlphanumeric(s string) bool {
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9' {
			continue
		}
		return false
	}
	return true
}

func IsNumeric(s string) bool {
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			continue
		}
		return false
	}
	return true
}

func IsLower(s string) bool {
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			continue
		}
		return false
	}
	return true
}
