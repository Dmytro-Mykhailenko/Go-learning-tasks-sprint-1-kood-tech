package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {

	var minI, maxI int

	if (from < 0 && to < 0) || (from >= len(arr) && to >= len(arr)) {
		return arr
	}
	if (from < 0 && to < 0) || (from >= len(arr) && to >= len(arr)) {
		return arr
	} else if from <= to {
		if from < 0 {
			minI = 0
		} else {
			minI = from
		}
		maxI = to
	} else if from >= to {
		if to < 0 {
			minI = 0
		} else {
			minI = to
		}
		maxI = from
	}
	arr = append(arr[:minI], arr[maxI:]...)

	return arr
}
