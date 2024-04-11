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

		if s[n] == ' ' {

			return 0
		}

		if (s[n] == '0' || s[n] == '-' || s[n] == '+') && isFirstZero {

			continue

		} else {

			isFirstZero = false
		}

		if !isFirstZero {

			x := ((len(s) - 1) - n)

			var mnoj = 1

			for i := 0; i < x; i++ {

				mnoj *= 10

			}

			resInt += int(s[n]-'0') * mnoj
		}
	}

	if isNegative {
		resInt = -resInt
	}

	return resInt
}
