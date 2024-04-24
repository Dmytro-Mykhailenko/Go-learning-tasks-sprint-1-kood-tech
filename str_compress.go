package sprint

import "fmt"

func StrCompress(input string) string {

	if len(input) < 2 {
		return input
	}

	sStr := []rune(input)
	var res string
	cnt := 1

	for i := 0; i < len(sStr)-1; i++ {

		if sStr[i] == sStr[i+1] {

			cnt += 1

			if i == len(sStr)-2 {
				res = res + fmt.Sprint(cnt) + string(sStr[i])
			}

			continue
		}

		if cnt > 1 {

			res = res + fmt.Sprint(cnt) + string(sStr[i])
			cnt = 1
			continue
		}

		if i == len(sStr)-2 {

			res = res + string(sStr[i]) + string(sStr[i+1])
			continue
		}

		res = res + string(sStr[i])
	}

	return res
}
