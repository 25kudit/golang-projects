package vehicles 

type Van struct {
	Vehicle
}

func NewVan(licensePlate string) (*Van) {
	return &Van{Vehicle: *NewVehicle(licensePlate, VehicleType(VAN))}
}