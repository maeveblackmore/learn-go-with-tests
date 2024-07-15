package structsmethodsinterfaces

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea) // `%#v` prints the value in Go-syntax format.
			}
		})
	}
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{32.0, 24.0}, hasPerimeter: 112.0},
		{name: "Circle", shape: Circle{12.0}, hasPerimeter: 75.39822368615503},
		{name: "Triangle", shape: Triangle{20.0, 40.0}, hasPerimeter: 100.0},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()

			if got != tt.hasPerimeter {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerimeter)
			}
		})
	}
}
