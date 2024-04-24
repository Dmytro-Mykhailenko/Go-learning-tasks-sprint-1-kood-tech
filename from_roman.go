package sprint

func FromRoman(s string) int {

	var x int

	for _, ch := range s {

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
	}

	return x
}
