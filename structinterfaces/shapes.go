package structinterfaces

import "math"

//Rectangle have hight and width.
type Rectangle struct {
	Width  float64
	Height float64
}

//Area calculates area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

//Circle has one property radius.
type Circle struct {
	Radius float64
}

//Area calculates area of the circle base on radius
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Triangle have lenght of base and hight as properties
type Triangle struct {
	Base   float64
	Height float64
}

//Area calculates area of the triangle based on base and hight
func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

//Perimeter calculates "perimeter" of the rectangle.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}
