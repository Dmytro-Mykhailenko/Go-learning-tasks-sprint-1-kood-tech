package sprint

func NbrBase(n int, base string) string {

	var resStr string
	var isNegative bool

	if n < 0 {

		isNegative = true
		n = -n
	}

	for n > 0 {

		x := n % 10

		resStr = string('0'+x) + resStr

		n /= 10
	}

	if isNegative {

		resStr = "-" + resStr
	}

	return resStr + " " + base
}
