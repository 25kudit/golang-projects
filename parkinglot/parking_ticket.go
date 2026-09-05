package main

import (
	"parkinglot/vehicles"
	"time"
)

const (
	BASE_CHARGES = 50
)

type ParkingTicket struct {
	EntryTime time.Time
	ExitTime time.Time
	Vehicle vehicles.VehicleInterface
	Spot *ParkingSpot
}

func NewParkingTicket(vehicle vehicles.VehicleInterface, spot *ParkingSpot) *ParkingTicket {
	return &ParkingTicket{EntryTime: time.Now(), ExitTime: time.Time{}, Vehicle: vehicle, Spot: spot}
}

func (pt *ParkingTicket) SetExitTime(exitTime time.Time) {
	pt.ExitTime = exitTime
}

func (pt *ParkingTicket) CalulateCharges() float64 {
	if time.Time.Equal(pt.ExitTime, time.Time{}) {
		return BASE_CHARGES
	}
	parkTime := pt.ExitTime.Sub(pt.EntryTime)
	charges := BASE_CHARGES + parkTime.Hours() * pt.Vehicle.GetVehicleCost()
	return charges
}