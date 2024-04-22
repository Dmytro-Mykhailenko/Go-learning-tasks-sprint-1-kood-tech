package sprint

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
