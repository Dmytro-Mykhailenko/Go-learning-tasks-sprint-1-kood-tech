package sprint

func NbrBase(n int, base string) string {

	bs := []rune(base)

	if len(bs) <= 2 {
		if len(bs) < 2 {
			return "NV"
		} else {
			if bs[0] == bs[1] {
				return "NV"
			}
		}
	}

	for i := 0; i < len(bs); i++ {
		cnt := 0

		for j := i + 1; j < len(bs); j++ {
			if bs[i] == '+' || bs[j] == '+' || bs[i] == '-' || bs[j] == '-' {
				return "NV"
			}

			if bs[i] == bs[j] {
				cnt++
				continue
			}
		}

		if cnt == len(bs)-1 {
			return "NV"
		}
	}
	return Itoa(n, base)
}

func Itoa(n int, base string) string {
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
