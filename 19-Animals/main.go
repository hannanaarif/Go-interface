package main

import "fmt"

// Animal interface: all animals must make a sound
type Animal interface {
	MakeSound()
}

// Animal registry: this stores functions (constructors) that create animals
var animalRegistry = make(map[string]func() Animal)

// RegisterAnimal function: registers an animal constructor in the registry
func RegisterAnimal(animalType string, constructor func() Animal) {
	animalRegistry[animalType] = constructor
}

// Dog type implements Animal
type Dog struct{}

func (d Dog) MakeSound() {
	fmt.Println("Woof!")
}

// AnimalFactory function: creates an animal from the registry
func AnimalFactory(animalType string) (Animal, error) {
	constructor, exists := animalRegistry[animalType]
	if !exists {
		return nil, fmt.Errorf("animal type '%s' not registered", animalType)
	}
	return constructor(), nil
}

func main() {
	// Register animals and their constructors
	RegisterAnimal("dog", func() Animal { return Dog{} })

	// Simulate creating animals dynamically
	animalTypes := []string{"dog", "cat"} // "cat" is not registered

	// Create and use the animals
	for _, animalType := range animalTypes {
		fmt.Printf("\nCreating animal type: %s\n", animalType)

		animal, err := AnimalFactory(animalType)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}
		animal.MakeSound()
	}
}
