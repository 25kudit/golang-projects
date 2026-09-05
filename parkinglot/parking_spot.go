package main

import (
	"fmt"
	"parkinglot/vehicles"
	"sync"
)

type ParkingSpot struct {
	ParkedVehicle vehicles.VehicleInterface
	VehicleType vehicles.VehicleType
	SpotId int 
	lock sync.Mutex
}

func NewParkingSpot(spotId int, vehicleType vehicles.VehicleType) (*ParkingSpot) {
	return &ParkingSpot{SpotId: spotId, VehicleType: vehicleType}
}

func (ps *ParkingSpot) IsAvailable() bool {
	return ps.ParkedVehicle == nil
}

func (ps *ParkingSpot) ParkVehicle(vehicle vehicles.VehicleInterface) error {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	if vehicle.GetVehicleType() != ps.VehicleType {
		return fmt.Errorf("vehicle type mismatch")
	}

	if ps.ParkedVehicle != nil {
		return fmt.Errorf("spot already occupied")
	}

	ps.ParkedVehicle = vehicle 
	return nil 
}

func (ps *ParkingSpot) UnparkVehicle() {
	ps.ParkedVehicle = nil
}