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
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{RectAngle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, v := range areaTests {
		got := v.shape.Area()

		if got != v.want {
			t.Errorf("got %.2f want %.2f", got, v.want)
		}
	}
}
