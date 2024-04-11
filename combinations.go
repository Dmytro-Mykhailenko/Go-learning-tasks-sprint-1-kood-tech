package sprint

import "fmt"

func Combinations() string {

	var resStr string

	for i := 0; i < 8; i++ {
		for j := i + 1; j < 9; j++ {
			for k := j + 1; k < 10; k++ {

				resStr += fmt.Sprintf("%d%d%d, ", i, j, k)

			}
		}
	}

	if len(resStr) > 2 {
		resStr = resStr[:len(resStr)-2]
	}

	return resStr

}
