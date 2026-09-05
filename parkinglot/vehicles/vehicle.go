package vehicles 

type VehicleType int

const (
	MOTORCYCLE VehicleType = iota 
	CAR 
	TRUCK
	VAN
)

var VehicleTypeMap = map[VehicleType]string {
	MOTORCYCLE : "MOTORCYCLE",
	CAR : "CAR",
	TRUCK : "TRUCK",
	VAN : "VAN",
}

var vehicleCosts = map[VehicleType]float64 {
	CAR : 100,
	TRUCK: 200,
	MOTORCYCLE: 50,
	VAN: 150,
}

type Vehicle struct {
	LicensePlate string
	VehicleType VehicleType
	Cost float64
}

type VehicleInterface interface {
	GetLiscensePlate() string
	GetVehicleType() VehicleType
	GetVehicleCost() float64
}

func (v *Vehicle) GetLiscensePlate() string {
	return v.LicensePlate
}

func (v *Vehicle) GetVehicleType() VehicleType {
	return v.VehicleType
}

func (v *Vehicle) GetVehicleCost() float64 {
	return v.Cost
}

func NewVehicle(licensePlate string, vehicleType VehicleType) (*Vehicle) {
	cost := vehicleCosts[vehicleType]
	return &Vehicle{LicensePlate: licensePlate, VehicleType: vehicleType, Cost: cost}
}