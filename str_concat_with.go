package sprint

func StrConcatWith(strs []string, sep string) string {
	var resStr string

	for i, str := range strs {

		if i != len(strs)-1 {
			resStr += str + sep
		} else {
			resStr += str
		}
	}
	return resStr
}
