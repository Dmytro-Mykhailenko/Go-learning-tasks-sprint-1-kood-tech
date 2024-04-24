package sprint

func LongestClimb(arr []int) []int {

	var in []int
	out := arr[:1]
	cnt := 1

	for i := 1; i < len(arr); i++ {

		if arr[i] > arr[i-1] && i != len(arr)-1 {

			in = append(in, arr[i-1])
			cnt++

			continue

		} else if arr[i] > arr[i-1] && i == len(arr)-1 {

			in = append(in, arr[i-1:]...)
			cnt++

		} else if cnt > 1 {

			in = append(in, arr[i-1])
		}

		if cnt > len(out) {

			out = in
		}

		in = []int{}
		cnt = 1
	}

	return out
}
