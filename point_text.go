package sprint

import "fmt"

type Point struct {
	X    float32
	Y    float32
	Text string
}

func PointText(p Point) Point {

	nP := MakePoint(p.X, p.Y, fmt.Sprintf("Text at (%v, %v)", p.X, p.Y))
	return nP
}
