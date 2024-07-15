package structsmethodsinterfaces

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{32.0, 24.0}, 112.0},
		{Circle{12.0}, 75.39822368615503},
		{Triangle{20.0, 40.0}, 100.0},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()

		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
