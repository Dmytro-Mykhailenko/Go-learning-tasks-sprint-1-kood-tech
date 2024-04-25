package sprint

func LongestCommonSubstr(str1, str2 string) string {

	res, lcs := "", ""

	for i := 0; i < len(str1); i++ {

		k := 0

		for j := 0; j < len(str2); j++ {

			if i+k > len(str1)-1 {
				break
			}

			if str1[i+k] == str2[j] {

				res += string(str2[j])
				k++

				continue
			} else if len(res) > len(lcs) {

				lcs = res

			}
			k = 0
			res = ""
		}

		if len(res) > len(lcs) {

			lcs = res
		}

		res = ""
	}

	return lcs
}
