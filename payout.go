package sprint

func Payout(amount int, denominations []int) (payout []int) {

	denominations = SortIntegerTable(denominations)

	var out []int
	i := len(denominations) - 1

	for {

		if amount >= denominations[i] {

			amount -= denominations[i]
			out = append(out, denominations[i])

		} else if i > 0 {

			i--
			continue
		}

		if amount == 0 {

			return out
		}

		if i == 0 && amount-denominations[i] == 0 {

			continue
		} else if i == 0 && amount-denominations[i] < 0 {
			out = []int{}
			return out
		}
	}
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
