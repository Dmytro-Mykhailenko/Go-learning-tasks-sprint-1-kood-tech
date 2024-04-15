package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {

	var minI, maxI int

	if from > to {
		minI, maxI = to, from
	}

	if to > from {

		minI, maxI = from, to
	}

	if minI < 0 {
		minI = 0
	}

	if maxI > len(arr) {
		maxI = len(arr)
	}

	arr = append(arr[:minI], arr[maxI:]...)

	return arr
}
