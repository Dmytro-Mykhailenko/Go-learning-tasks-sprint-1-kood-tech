package sprint

func Countdown(n int) string {

	var resStr string

	for i := n; i > 0; i -= 2 {

		resStr += string('0'+i) + ", "

	}

	resStr += "0!"

	return resStr

}
