package main

import "fmt"

// Vehicle interface: all vehicles must be able to drive
type Vehicle interface {
	Drive() // Every vehicle must implement the Drive method
}

// Vehicle registry: stores constructors for creating vehicle instances
var vehicleRegistry = make(map[string]func() Vehicle)

// RegisterVehicle registers a constructor for a specific vehicle type
func RegisterVehicle(vehicleType string, constructor func() Vehicle) {
	vehicleRegistry[vehicleType] = constructor
}

// Car type implements Vehicle
type Car struct{}

func (c Car) Drive() {
	fmt.Println("The car drives on the road!")
}

// Bike type implements Vehicle
type Bike struct{}

func (b Bike) Drive() {
	fmt.Println("The bike rides on the road!")
}

// Truck type implements Vehicle
type Truck struct{}

func (t Truck) Drive() {
	fmt.Println("The truck drives on the highway!")
}

// VehicleFactory creates a vehicle dynamically based on the type
func VehicleFactory(vehicleType string) (Vehicle, error) {
	constructor, exists := vehicleRegistry[vehicleType]
	if !exists {
		return nil, fmt.Errorf("vehicle type '%s' not registered", vehicleType)
	}
	return constructor(), nil
}

func main() {
	// Register different vehicles and their constructor functions
	RegisterVehicle("car", func() Vehicle { return Car{} })
	RegisterVehicle("bike", func() Vehicle { return Bike{} })
	RegisterVehicle("truck", func() Vehicle { return Truck{} })

	// Simulate creating and driving vehicles
	vehicleTypes := []string{"car", "bike", "truck", "bus"} // 'bus' is not registered

	for _, vehicleType := range vehicleTypes {
		fmt.Printf("\nCreating vehicle type: %s\n", vehicleType)

		vehicle, err := VehicleFactory(vehicleType)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}
		vehicle.Drive()
	}
}
