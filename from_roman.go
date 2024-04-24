package sprint

func FromRoman(s string) int {

	var x int
	var res int

	for x < len(s) {

		if x < len(s)-1 {

			if retRomNum(rune(s[x])) < retRomNum(rune(s[x+1])) {

				res -= retRomNum(rune(s[x]))
			} else {

				res += retRomNum(rune(s[x]))
			}

		} else {

			res += retRomNum(rune(s[x]))
		}

		x++
	}

	return res
}

func retRomNum(ch rune) int {

	var x int

	switch ch {

	case 'M':
		x += 1000

	case 'D':
		x += 500

	case 'C':
		x += 100

	case 'L':
		x += 50

	case 'X':
		x += 10

	case 'V':
		x += 5

	case 'I':
		x += 1
	}

	return x
}
