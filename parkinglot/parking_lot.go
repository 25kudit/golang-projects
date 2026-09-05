package main

import (
	"fmt"
	"parkinglot/vehicles"
	"sync"
	"time"
)	

type ParkingLot struct {
	Name string
	floors []*ParkLevel
}

var (
	ParkingLotInstance *ParkingLot
	once sync.Once
)

func GetParkingLotInstance() *ParkingLot {
	once.Do(func() {
		fmt.Println("Creating new parking lot instance")
		ParkingLotInstance = &ParkingLot{}
	})
	return ParkingLotInstance
}

func (p *ParkingLot) AddParkLevel(floorId int) {
	p.floors = append(p.floors, NewParkLevel(floorId))
} 

func (p *ParkingLot) findParkingSpot(vehicleType vehicles.VehicleType) (*ParkingSpot, error) {
	for _, floor := range p.floors {
		if spot := floor.FindParkingSpot(vehicleType); spot != nil {
			return spot, nil 
		}
	}
	return nil, fmt.Errorf("No available parking spot for %s", vehicles.VehicleTypeMap[vehicleType])
} 

func (p *ParkingLot) ParkVehicle(vehicle vehicles.VehicleInterface) (*ParkingTicket, error) {
	parkingSpot, err := p.findParkingSpot(vehicle.GetVehicleType())
	if err != nil {
		return nil, err
	}

	err = parkingSpot.ParkVehicle(vehicle)
	if err != nil {
		return nil, err
	}
	parkingTicket := NewParkingTicket(vehicle, parkingSpot)
	return parkingTicket, nil
}

func (p *ParkingLot) UnparkVehicle(parkingTicket *ParkingTicket) float64 {
	parkingTicket.SetExitTime(time.Now())
	totalCharges := parkingTicket.CalulateCharges()
	parkingTicket.Spot.UnparkVehicle()
	return totalCharges
}

func (p *ParkingLot) DisplayAvailability() {
	for _, floor := range p.floors {
		floor.DisplayAvailability()
	}
}