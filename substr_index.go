package sprint

func SubstrIndex(s string, toFind string) int {

	if len(toFind) == 0 || toFind == s {

		return 0
	}

	cnt := 0

	for i := 0; i <= len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind {
			cnt++
		}
	}

	if cnt == 0 {
		return -1
	}

	return cnt
}
