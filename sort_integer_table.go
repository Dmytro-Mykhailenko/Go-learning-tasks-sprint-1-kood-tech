package sprint

func SortIntegerTable(table []int) []int {

	// x := table[0]
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
