package sprint

func StrReverse(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}
