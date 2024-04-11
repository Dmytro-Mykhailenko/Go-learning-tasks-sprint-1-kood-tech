package sprint

func AlphaNumber(n int) string {

	var resStr string
	var isNegative bool

	if n < 0 {

		isNegative = true
		n = -n
	}

	for n > 0 {

		x := n % 10

		resStr = string('a'+x) + resStr

		n /= 10
	}

	if isNegative {

		resStr = "-" + resStr
	}

	return resStr

}
