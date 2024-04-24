package sprint

func RemoveDuplicates(arr []int) []int {

	var unique = make(map[int]int)
	var out []int

	for _, v := range arr {
		if _, ok := unique[v]; ok {
			continue
		}

		unique[v] = v
		out = append(out, v)
	}

	return out
}
