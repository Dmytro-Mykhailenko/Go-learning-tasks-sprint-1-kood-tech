package sprint

func StrCompare(a, b string) int {

	aSlice := []rune(a)
	bSlice := []rune(b)

	var longerSlc, shorterSlc []rune
	var aSlcLngr bool

	if a == b {
		return 0
	} else if len(aSlice) > len(bSlice) {
		longerSlc = aSlice
		shorterSlc = bSlice
		aSlcLngr = true
	} else {
		longerSlc = bSlice
		shorterSlc = aSlice
	}

	for i := 0; i < len(shorterSlc); i++ {
		if longerSlc[i] == shorterSlc[i] {
			if i == len(shorterSlc)-1 {
				if aSlcLngr {
					return 1
				} else {
					return -1
				}
			}
			continue
		} else if aSlice[i] < bSlice[i] {
			return -1
		} else {
			return 1
		}
	}

	return 0
}
