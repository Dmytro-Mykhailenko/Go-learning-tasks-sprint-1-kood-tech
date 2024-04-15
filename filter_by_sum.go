package sprint

func FilterBySum(arr [][]int, limit int) [][]int {

	for i := 0; i < len(arr); i++ {

		var sum int

		for j := 0; j < len(arr[i]); j++ {
			sum += arr[i][j]
		}

		if sum < limit {
			arr = append(arr[:i], arr[i+1:]...)
			i -= 1
		}
	}
	return arr
}
