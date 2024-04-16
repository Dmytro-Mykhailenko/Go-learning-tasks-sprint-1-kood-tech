package sprint

func SubstrIndex(s string, toFind string) int {

	if len(toFind) == 0 || toFind == s {

		return 0
	}

	cnt := 0

	spr := []rune(toFind)
	str := []rune(s)

	for i := 0; i <= len(str)-len(spr); i++ {
		for j := 0; j < len(spr); j++ {
			if str[i+j] == spr[j] && j == len(spr)-1 {
				cnt++
				i += j
			} else if str[i+j] == spr[j] && j != len(spr)-1 {
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
