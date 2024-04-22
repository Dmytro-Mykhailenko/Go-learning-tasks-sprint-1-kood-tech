package sprint

func NthFibonacci(index int) int {

	if index < 0 {
		return -1
	}
	
	if index == 0 {
		return 0
	}

	if index == 1 {
		return 1
	}

	return NthFibonacci(index-2) + NthFibonacci(index-1)
}