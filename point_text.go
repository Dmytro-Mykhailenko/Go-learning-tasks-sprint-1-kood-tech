package sprint

import "fmt"

type Point struct {
	X    float32
	Y    float32
	Text string
}

func PointText(p Point) Point {

	nP := MakePoint(p.X, p.Y, fmt.Sprintf("Text at (%f, %f)", p.X, p.Y))
	return nP
}

func MakePoint(x, y float32, text string) Point {

	var p Point = Point{x, y, text}

	return p
}
