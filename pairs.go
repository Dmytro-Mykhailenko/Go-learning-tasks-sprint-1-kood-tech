package sprint

import "fmt"

func Pairs() string {

	var resStr, s1, s2 string

	for i := 0; i < 99; i++ {

		if i < 10 {

			s1 = "0" + fmt.Sprint(i)

		} else {

			s1 = fmt.Sprint(i)

		}

		for j := i + 1; j < 100; j++ {

			if j < 10 {

				s2 = "0" + fmt.Sprint(j)

			} else {

				s2 = fmt.Sprint(j)

			}

			resStr += s1 + " " + s2 + ", "

		}

	}

	return resStr

}
