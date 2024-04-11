package sprint

func StrToInt(s string) int {

	var resInt int

	isFirstZero := true
	var isNegative bool

	if (s[0] == '-' && s[1] == '-') || (s[0] == '+' && s[1] == '+') {

		return 0

	} else if s[0] == '-' && s[1] != '-' {

		isNegative = true
	}

	for n := 0; n < len(s); n++ {

		if (s[n] == '0' || s[n] == '-' || s[n] == '+') && isFirstZero {

			continue

		} else if s[n] < '0' || s[n] > '9' {

			return 0
		} else {
			isFirstZero = false
		}

		if !isFirstZero {

			x := ((len(s) - 1) - n)

			var val = 1

			for i := 0; i < x; i++ {

				val *= 10

			}

			resInt += int(s[n]-'0') * val
		}
	}

	if isNegative {
		resInt = -resInt
	}

	return resInt
}
