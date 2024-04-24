package sprint

import (
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {

	sStr1 := strings.Split(strings.ToLower(str1), "")
	sStr2 := strings.Split(strings.ToLower(str2), "")

	sort.Strings(sStr1)
	sort.Strings(sStr2)

	m := strings.TrimSpace(strings.Join(sStr1, ""))
	s := strings.TrimSpace(strings.Join(sStr2, ""))

	return strings.EqualFold(m, s)
}
