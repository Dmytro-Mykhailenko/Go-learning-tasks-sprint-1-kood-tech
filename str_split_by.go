package sprint

func StrSplitBy(s, sep string) []string {
	res := []string{}
	var word, spr []byte

	var IsSep bool
	cnt := 0

	for i := 0; i < len(s); i++ {

		if s[i] == sep[cnt] {
			if cnt == len(sep)-1 {
				IsSep = true
				spr = []byte{}
			} else {
				spr = append(spr, s[i])
				cnt++
				continue
			}
		} else {
			word = append(spr, word...)
			spr = []byte{}
			cnt = 0
		}

		if IsSep {
			res = append(res, string(spr)+string(word))
			word = []byte{}
			IsSep = false
			cnt = 0
			continue
		}

		word = append(word, s[i])

		if i == len(s)-1 && !IsSep {
			res = append(res, string(word))
		}
	}

	return res
}
