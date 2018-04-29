package shapes

import "math"

type RectAngle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (r RectAngle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle RectAngle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle RectAngle) float64 {
	return rectangle.Width * rectangle.Height
}
