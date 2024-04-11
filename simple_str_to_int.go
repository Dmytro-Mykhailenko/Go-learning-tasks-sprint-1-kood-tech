package sprint

func SimpleStrToInt(s string) int {

	var resInt int

	isFirstZero := true

	for n := 0; n < len(s); n++ {

		if s[n] == '-' || s[n] == '+' || s[n] == ' ' {

			return 0
		}

		if s[n] == '0' && isFirstZero {

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

	return resInt
}
