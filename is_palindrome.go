package sprint

func IsPalindrome(s string) bool {

	str := []rune(s)

	for i, ch := range str {
		if ch != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
