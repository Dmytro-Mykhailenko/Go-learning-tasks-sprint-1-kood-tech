package spring

func Accumulate(n int) int {

	var sum int

	if n > 0 {

		for i := 1; i <= n; i++ {

			sum += i

		}
		return sum
	}
	return 0
}
