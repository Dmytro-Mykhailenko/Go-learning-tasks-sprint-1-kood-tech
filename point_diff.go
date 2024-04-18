package sprint

func PointDiff(p1, p2 Point) Point {

	if p1.X >= p2.X {
		return p1
	}
	return p2
}
