package sprint

import "unicode/utf8"

func StrLength(s string) []int {
	return []int{utf8.RuneCountInString(s), len(s)}
}
