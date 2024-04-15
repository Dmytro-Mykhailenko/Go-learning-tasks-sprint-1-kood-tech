package sprint

func BalanceOut(arr []bool) []bool {

	var count int
	var element bool

	for _, x := range arr {
		if x {
			count++
		} else {
			count--
		}
	}

	if count == 0 {
		return arr
	} else if count < 0 {
		count = -count
		element = true
	}

	for ; 0 < count; count-- {
		arr = append(arr, element)
	}

	return arr
}
