package main

type Shape interface {
	Area()

	Perimeter()
}

type Rectangle struct {
}

type Circle struct {
}

func (r *Rectangle) Area() {
	println("Rectangle.Area")
}

func (c *Circle) Area() {
	println("Circle.Area")
}

func (r *Rectangle) Perimeter() {
	println("Rectangle.Perimeter")
}

func (c *Circle) Perimeter() {
	println("Circle.Perimeter")
}

func main() {

	var s Shape = &Rectangle{}
	s.Area()
	s.Perimeter()

	var c Shape = &Circle{}
	c.Area()
	c.Perimeter()
}
