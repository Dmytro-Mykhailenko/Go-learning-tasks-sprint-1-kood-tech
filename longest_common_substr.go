package sprint

func LongestCommonSubstr(str1, str2 string) string {

	found := false
	res := ""
	lcs := ""

	for i := 0; i < len(str1); i++ {

		k := 0

		for j := 0; j < len(str2); j++ {

			if str1[i+k] == str2[j] {

				found = true
				res += string(str2[j])
				k++
				if i+k > len(str1)-1 {
					break
				}
				// continue
			} else if found {

				break
			}
		}

		if found && len(res) > len(lcs) {

			lcs = res
		}

		found = false
		res = ""
	}

	return lcs
}
