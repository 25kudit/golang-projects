package vehicles

type Car struct {
	Vehicle
}

func NewCar(licensePlate string) (*Car) {
	return &Car{Vehicle: *NewVehicle(licensePlate, VehicleType(CAR))}
}