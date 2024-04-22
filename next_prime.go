package sprint

func NextPrime(n int) int {

	var isprime bool

	for i := n; ; i++ {

		if i < 2 {
			continue
		}

		isprime = true

		for j := 2; j <= i/j; j++ {

			if i%j == 0 {
				isprime = false
			}
		}

		if isprime {

			return i
		}
	}
}
