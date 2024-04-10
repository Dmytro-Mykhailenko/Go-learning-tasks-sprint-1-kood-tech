package sprint

func IsLeapYear(year int) bool {

	if year%400 != 0 {
		return false
	} else if year%4 == 0 {
		return true
	} else {
		return false
	}
}
