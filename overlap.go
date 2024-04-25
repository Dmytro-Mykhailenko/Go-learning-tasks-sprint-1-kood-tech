package sprint

func Overlap(arr1, arr2 []int) []int {

	arr1 = SortIntegerTable(arr1)
	arr2 = SortIntegerTable(arr2)

	out := []int{}
	found := false
	j := 0
	f := 0
	for i := 0; i < len(arr1); i++ {

		for ; j < len(arr2); j++ {

			if arr1[i] == arr2[j] {
				found = true
				break
			}
		}

		if found {
			out = append(out, arr1[i])
			found = false
			if j != len(arr2)-1 {
				j++
			}
			f = j
		} else {
			j = f
		}
	}

	return out
}

func SortIntegerTable(table []int) []int {

	var x int

	for i := 0; i < len(table)-1; i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				x = table[i]
				table[i] = table[j]
				table[j] = x
			}
		}
	}
	return table
}
