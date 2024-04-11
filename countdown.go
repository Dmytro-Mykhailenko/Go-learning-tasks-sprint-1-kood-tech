package sprint

import "fmt"

func Countdown(n int) string {

	var resStr string

	for i := n; n > 0; i -= 2 {

		resStr += fmt.Sprint(i) + ", "

	}

	resStr += ", 0!"

	return resStr

}
