package sprint

func TransposeMatrix(matrix [][]int) [][]int {

	out := [][]int{}

	for i := 0; i < len(matrix[0]); i++ {

		out = append(out, []int{})
		for j := 0; j < len(matrix); j++ {

			out[i] = append(out[i], matrix[j][i])
		}
	}

	return out
}
