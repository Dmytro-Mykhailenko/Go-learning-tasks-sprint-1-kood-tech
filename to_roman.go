package sprint

func ToRoman(num int) string {

	if num > 3999 || num < 1 {

		return "Invalid input"
	}

	b := []int{1000, 500, 100, 50, 10, 5, 1}
	out := ""
	i := 0

	for {

		if num >= b[i] {

			num -= b[i]
			out += retNumRom(b[i])

		}

		if num == 0 {

			break

		}

		if num < b[i] {

			if i < len(b)-2 {

				if num < b[i] && b[i]-num <= b[i+2] {

					num += b[i+2]
					out += retNumRom(b[i+2])
					continue
				}
			}
			if i < len(b)-1 {

				i++
				continue
			}
		}
	}

	return out
}

func retNumRom(z int) string {

	var x string

	switch z {

	case 1000:
		x += "M"

	case 500:
		x += "D"

	case 100:
		x += "C"

	case 50:
		x += "L"

	case 10:
		x += "X"

	case 5:
		x += "V"

	case 1:
		x += "I"
	}

	return x
}
