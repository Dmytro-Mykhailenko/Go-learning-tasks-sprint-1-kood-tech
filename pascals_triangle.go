package sprint

func PascalsTriangle(n int) [][]int {

	var out [][]int

	for i := 0; i < n; i++ {

		out = append(out, []int{})

		for j := 0; j <= i; j++ {

			if j == 0 || i-j == 0 {

				out[i] = append(out[i], 1)

			} else if j == 1 || i-j == 1 {

				out[i] = append(out[i], i)
			} else {

				out[i] = append(out[i], out[i-1][j-1]+out[i-1][j])
			}
		}
	}

	return out
}
