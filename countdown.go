package sprint

import "fmt"

func Countdown(n int) string {

	var resStr string

	for n; n > 0; n -= 2 {

		resStr += fmt.Sprint(n) + ", "

	}

	resStr += ", 0!"

	return resStr

}
