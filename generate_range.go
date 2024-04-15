package sprint

func GenerateRange(min, max int) []int {

	var arrayOfInt []int

	if min >= max {
		return arrayOfInt
	}

	for i := min; i < max; i++ {
		arrayOfInt = append(arrayOfInt, i)
	}
	return arrayOfInt
}
