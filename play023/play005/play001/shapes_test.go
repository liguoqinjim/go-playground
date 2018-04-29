package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := RectAngle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %0.2f want %0.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := RectAngle{12.0, 6.0}
	got := Area(rectangle)
	want := 72.0

	if got != want {
		t.Errorf("got %0.2f want %0.2f", got, want)
	}
}
