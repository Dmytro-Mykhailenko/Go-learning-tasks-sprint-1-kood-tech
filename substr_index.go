package sprint

func SubstrIndex(s string, toFind string) int {

	if len(toFind) == 0 || toFind == s {

		return 0
	}

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
