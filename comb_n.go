package sprint

func CombN(n int) []string {

	var combStr []string

	prefix := ""

	for i := 0; i < n; i++ {
		prefix = prefix + string(i+48)
	}

	combStr = append(combStr, prefix)

	slicePrefix := []rune(prefix)

	x := len(slicePrefix) - 1

	for x >= 0 {

		y := int(slicePrefix[x]) - 48

		for y < 10-n+x {

			y++

			slicePrefix[x] = rune(y) + 48
			combStr = append(combStr, string(slicePrefix))

			if x < len(slicePrefix)-1 {
				if slicePrefix[x] < slicePrefix[x+1] {

					x++
					combStr = combStr[:len(combStr)-1]
				}
			}
		}
		x--
	}

	return combStr
}
