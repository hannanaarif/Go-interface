package main
import "fmt"


type Animal interface {
	Speak() string
}

type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return "Woof!"
}


type Cat struct {
	name string
}

func (c Cat) Speak() string {
	return "Meow!"
}


func main() {
	dog:= Dog{name: "Fido"}

	cat:= Cat{name: "Whiskers"}

	var d Animal= dog
	var c Animal= cat

	fmt.Println(d.Speak())
	fmt.Println(c.Speak())
}