package structsmethodsinterfaces

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// A method is a function with a receiver, using the syntax `func (receiverName ReceiverType) MethodName(args)`.
// In this case, `(r Rectangle)` is our receiver, which means the method `Area()` is associated with the Rectangle type.
// The receiver name `r` allows the method to access the fields of that `Rectangle` instance.

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}
