package sprint

func IsSorted(f func(a, b string) int, arr []string) bool {
	var sortAsc, sortDesc bool

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if f(arr[i], arr[j]) > 0 {
				if sortAsc {
					return false
				}
				sortDesc = true
			} else if f(arr[i], arr[j]) < 0 {
				if sortDesc {
					return false
				}
				sortAsc = true
			}
		}
	}

	return true
}

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
