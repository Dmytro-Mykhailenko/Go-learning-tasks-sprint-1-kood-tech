package sprint

func BetweenLimits(from, to rune) string {

	var finalStr string
	from += 1

	for from < to {

		finalStr += string(from)
		from += 1

	}
	return finalStr
}
