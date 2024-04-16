package sprint

func IsNumeric(s string) bool {
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			continue
		}
		return false
	}
	return true
}
