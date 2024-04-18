package sprint

type Point struct {
	X    float32
	Y    float32
	Text string
}

func PointDiff(p1, p2 Point) Point {

	if p1.X >= p2.X {
		return p1
	}
	return p2
}
