package sprint

func Payout(amount int, denominations []int) (payout []int) {

	var out []int

	for i := len(denominations) - 1; i >= 0; i-- {

		if amount >= denominations[i] {

			amount -= denominations[i]
			out = append(out, denominations[i])

			if amount == 0 {

				return out
			}

			if amount < 0 {

				out = []int{}
				break
			}
		}
	}

	return out
}
