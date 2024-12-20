package main

import "fmt"

type Greeter interface {
	greet() string
}

type Person struct {
	Name string
}

func (p Person) greet() string {
	return "Hello, my name is " + p.Name
}

type Robot struct {
	ID int
}

func (r Robot) greet() string {
	return fmt.Sprintf("Beep Boop. I am Robot #%d", r.ID)
}

func sayHello(g Greeter) string {
	return g.greet()
}

func main() {
	p:= Person{Name: "Alice"}
	r:= Robot{ID: 42}

	fmt.Println(sayHello(p))
	fmt.Println(sayHello(r))
}