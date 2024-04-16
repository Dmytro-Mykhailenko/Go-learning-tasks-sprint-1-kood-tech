package sprint

func StrLength(s string) []int {
	runeCount := 0

	for range s {
		runeCount++
	}

	return []int{runeCount, len(s)}
}
