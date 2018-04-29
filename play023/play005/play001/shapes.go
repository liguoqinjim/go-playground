package shapes

type RectAngle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle RectAngle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle RectAngle) float64 {
	return rectangle.Width * rectangle.Height
}
