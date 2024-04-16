package sprint

func SubstrIndex(s string, toFind string) int {

	if len(toFind) == 0 || toFind == s {

		return 0
	}

	cnt := 0

	for i := 0; i <= len(s)-len(toFind); i++ {
		for j := 0; j < len(toFind); j++ {
			if s[i+j] == toFind[j] && j == len(toFind)-1 {
				cnt++
				i += j
			} else if s[i+j] == toFind[j] && j != len(toFind)-1 {
				continue
			} else {
				break
			}
		}
	}

	if cnt == 0 {
		return -1
	}

	return cnt
}
