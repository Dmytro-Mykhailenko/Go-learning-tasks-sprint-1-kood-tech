package sprint

func StrSplitBy(s, sep string) []string {
	res := []string{}
	var word []byte
	var IsSep bool
	i := 0

	for i < len(s) {
		if s[i] == sep[0] {
			for j := 0; j < len(sep); j++ {

				if s[i+j] != sep[j] {

					IsSep = false
					break
				} else if j == len(sep)-1 {
					IsSep = true
				}

			}

			if IsSep {
				res = append(res, string(word))
				word = []byte{}
				i += len(sep)
				IsSep = false
			}
		}
		word = append(word, s[i])
		if i == len(s)-1 && !IsSep {
			res = append(res, string(word))
		}
		i++
	}

	return res
}
