package main

import "fmt"

// Base interface
type Animal interface {
	MakeSound()
}

// Extended interface: embeds the base interface and adds more methods
type MovableAnimal interface {
	Animal
	Move()
}

// Define a Dog struct
type Dog struct{}

func (d Dog) MakeSound() {
	fmt.Println("Woof!")
}

func (d Dog) Move() {
	fmt.Println("The dog runs.")
}

// Define a Fish struct
type Fish struct{}

func (f Fish) MakeSound() {
	fmt.Println("Blub!")
}

func (f Fish) Move() {
	fmt.Println("The fish swims.")
}

// Define a Bird struct
type Bird struct{}

func (b Bird) MakeSound() {
	fmt.Println("Chirp!")
}

func (b Bird) Move() {
	fmt.Println("The bird flies.")
}

// Function to interact with any Animal
func Interact(animal Animal) {
	animal.MakeSound()

	// Use type assertion to check if the animal is also MovableAnimal
	if movable, ok := animal.(MovableAnimal); ok {
		fmt.Println("This animal can move!")
		movable.Move()
	} else {
		fmt.Println("This animal does not move visibly!")
	}
}

func main() {
	// Create instances
	dog := Dog{}
	fish := Fish{}
	bird := Bird{}

	// Interact with each animal
	fmt.Println("Interacting with the dog:")
	Interact(dog)
	fmt.Println("\nInteracting with the fish:")
	Interact(fish)
	fmt.Println("\nInteracting with the bird:")
	Interact(bird)
}
