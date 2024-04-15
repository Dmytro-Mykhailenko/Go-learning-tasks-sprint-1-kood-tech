package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {

	var minI, maxI int

	// if (from < 0 && to < 0) || (from >= len(arr) && to >= len(arr)) {
	// 	return arr
	if from < to {
		if f < 0 {
			minI = 0
		} else {
			minI = from
		}
		if to > len(arr) {
			maxI = len(arr)
		} else {
			maxI = to
		}
	}

	if from > to {
		if to < 0 {
			minI = 0
		} else {
			minI = to
		}

		if from > len(arr) {
			maxI = len(arr)
		} else {
			maxI = from
		}
	}

	arr = append(arr[:minI], arr[maxI:]...)

	return arr
}
