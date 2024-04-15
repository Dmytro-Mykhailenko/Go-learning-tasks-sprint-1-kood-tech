package sprint

func BulkAtoi(arr []string) []int {

	var arrInt []int

	for _, str := range arr {

		arrInt = append(arrInt, Atoi(str))
	}
	return arrInt
}

func Atoi(str string) int {

	var isNegative bool
	var isPositive bool
	var isNumber bool
	var res int

	for i := 0; i < len(str); i++ {
		if str[i] == '-' && !isNegative && !isPositive && !isNumber {
			isNegative = true
			continue
		} else if str[i] == '+' && !isPositive && !isNegative && !isNumber {
			isPositive = true
			continue
		} else if (str[i] == '-' || str[i] == '+') && (isNumber || isNegative || isPositive) {
			return 0
		} else if str[i] == '0' && !isNumber {
			continue
		} else if str[i] <= '9' && str[i] >= '0' {
			isNumber = true
			res += int(str[i]-48) * powerOf10(len(str)-i)
		} else if (str[i] < '0' || str[i] > '9') && (!isPositive && !isNegative && !isNumber) {
			continue
		} else {
			return 0
		}
	}

	if isNegative {
		res = -res
	}

	return res
}

func powerOf10(n int) int {
	if n == 1 {
		return 1
	}
	return powerOf10(n-1) * 10
}
