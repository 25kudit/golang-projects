package vehicles

type MotorCycle struct {
	Vehicle
}

func NewMotorCycle(licensePlate string) (*MotorCycle) {
	return &MotorCycle{Vehicle: *NewVehicle(licensePlate, VehicleType(MOTORCYCLE))}
}