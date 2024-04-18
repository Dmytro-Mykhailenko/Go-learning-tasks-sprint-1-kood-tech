package sprint

type Circle struct {
	Radius    float32
	Diameter  float32
	Area      float32
	Perimeter float32
}

func GetCircle(r float32) Circle {
	P := float32(3.14)

	var circle Circle

	circle.Radius = r
	circle.Diameter = 2*r
	circle.Area = P*(r*r)
	circle.Perimeter = 2*P*r
	
	return circle
}
