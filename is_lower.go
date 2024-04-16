package sprint

func IsLower(s string) bool {
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			continue
		}
		return false
	}
	return true
}
