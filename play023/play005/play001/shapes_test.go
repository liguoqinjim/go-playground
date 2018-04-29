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
	t.Run("rectangle", func(t *testing.T) {
		rectangle := RectAngle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %0.2f want %0.2f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.16
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
