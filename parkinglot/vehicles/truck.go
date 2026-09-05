package vehicles

type Truck struct {
	Vehicle
}

func NewTruck(licensePlate string) (*Truck) {
	return &Truck{Vehicle: *NewVehicle(licensePlate, VehicleType(TRUCK))}
}