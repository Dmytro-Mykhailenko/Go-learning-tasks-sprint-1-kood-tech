package sprint

func BalancedParentheses(input string) bool {

	var isRoundBracketsOpened, isCurlyBracesOpened, isSquareBracketsOpened bool
	var roundBracketsCount, squareBracketsCount, CurlyBracesCount int

	for _, ch := range input {

		switch ch {

		case '(':
			{
				isRoundBracketsOpened = true
				roundBracketsCount++
			}
		case '[':
			{
				isSquareBracketsOpened = true
				squareBracketsCount++
			}
		case '{':
			{
				isCurlyBracesOpened = true
				CurlyBracesCount++
			}
		case ')':
			{
				if !isRoundBracketsOpened {
					return false
				}

				roundBracketsCount--

				if roundBracketsCount == 0 {
					isRoundBracketsOpened = false
				}
			}
		case ']':
			{
				if !isSquareBracketsOpened {
					return false
				}

				squareBracketsCount--

				if squareBracketsCount == 0 {
					isSquareBracketsOpened = false
				}
			}
		case '}':
			{
				if !isCurlyBracesOpened {
					return false
				}

				CurlyBracesCount--

				if CurlyBracesCount == 0 {
					isCurlyBracesOpened = false
				}
			}
		}
	}

	if roundBracketsCount == 0 && squareBracketsCount == 0 && CurlyBracesCount == 0 {
		return true
	}

	return false
}
