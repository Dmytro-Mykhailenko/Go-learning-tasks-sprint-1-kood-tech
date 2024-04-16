package sprint

func SubstrIndex(s string, toFind string) int {

	cnt := 0
	for i, _ := range s {
		if i+len(toFind) <= len(s) {
			if s[i:i+len(toFind)] == toFind {
				cnt++
			}
		}
	}
	if cnt == 0 {
		return -1
	}
	return cnt
}
