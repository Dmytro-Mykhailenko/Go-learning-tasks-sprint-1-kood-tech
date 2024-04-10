package sprint

func IntVsFloat(i int, f float32) string {

	intStr := "Integer"
	fltStr := "Float"
	samStr := "Same"
	z := float32(i)

	if z > f {
		return intStr
	} else if f > z {
		return fltStr
	} else {
		return samStr
	}
}
