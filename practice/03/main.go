package main

import "fmt"

type Shape interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14159 * c.radius * c.radius
}

func getArea(s Shape) float64{
	return s.area()
}

func main(){
	r:=Rectangle{width: 5, height: 10}
	c:=Circle{radius: 5}
	fmt.Println("Rectangle area:", getArea(r))
	fmt.Println("Circle area:", getArea(c))
}