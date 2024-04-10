package sprint

func BetweenLimits(from, to rune) string {

	var finalStr string

	if from < to {
		from += 1
		for from < to {

			finalStr += string(from)
			from += 1

		}
	} else {

		to += 1
		for to < from {

			finalStr += string(to)
			to += 1

		}
	}

	return finalStr
}
