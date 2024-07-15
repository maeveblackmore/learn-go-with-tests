package structsmethodsinterfaces

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// A method is a function with a receiver, using the syntax `func (receiverName ReceiverType) MethodName(args)`.
// In this case, `(r Rectangle)` is our receiver, which means the method `Area()` is associated with the Rectangle type.
// The receiver name `r` allows the method to access the fields of that `Rectangle` instance.

// Returns the area of a rectangle given the width and height.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Returns the perimeter of a rectangle given the width and height.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

// Returns the area of a circle given the radius.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Returns the perimeter (circumference) of a circle given the radius.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Triangle struct {
	base   float64
	height float64
}

// Returns the area of an isosceles triangle given the base and height.
func (t Triangle) Area() float64 {
	return (t.base * t.height) / 2
}

// Returns the perimeter of an isosceles triangle given the base and height.
func (t Triangle) Perimeter() float64 {
	return t.base + (t.height * 2)
}
