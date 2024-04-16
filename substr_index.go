package sprint

func SubstrIndex(s string, toFind string) int {
	var j, cnt int
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[j] && j != len(toFind)-1 {
			j++
			continue
		} else if s[i] == toFind[j] && j == len(toFind)-1 {
			j = 0
			cnt += 1
		}
	}
	if cnt == 0 {
		return -1
	}
	return cnt
}
