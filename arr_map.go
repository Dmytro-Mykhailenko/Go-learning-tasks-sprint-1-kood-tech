package sprint

func ArrMap(f func(int) bool, a []int) []bool {
	var boolSlice []bool

	for _, el := range a {
		boolSlice = append(boolSlice, f(el))
	}

	return boolSlice
}

func IsNegative(n int) bool {

	return n < 0

}

func IsPrime(n int) bool {

	if n > 1 {

		for i := 2; i <= n/i; i++ {

			if n%i == 0 {

				return false
			}
		}

		return true
	}

	return false
}
