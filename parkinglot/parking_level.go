package main

import (
	"fmt"
	"parkinglot/vehicles"
)

var spotCount = map[vehicles.VehicleType]int {
	vehicles.MOTORCYCLE : 50,
	vehicles.CAR : 15,
	vehicles.TRUCK : 5,
	vehicles.VAN : 10,
}

type ParkLevel struct {
	FloorId int
	ParkingSpots map[vehicles.VehicleType]map[int]*ParkingSpot
}

func NewParkLevel(floorId int) *ParkLevel {
	parkingSpots := make(map[vehicles.VehicleType]map[int]*ParkingSpot)

	parkingSpots[vehicles.CAR] = createParkingSpots(vehicles.CAR)
	parkingSpots[vehicles.MOTORCYCLE] = createParkingSpots(vehicles.MOTORCYCLE)
	parkingSpots[vehicles.TRUCK] = createParkingSpots(vehicles.TRUCK)
	parkingSpots[vehicles.VAN] = createParkingSpots(vehicles.VAN)

	return &ParkLevel{FloorId: floorId, ParkingSpots: parkingSpots}
}

func createParkingSpots(vehicleType vehicles.VehicleType) map[int]*ParkingSpot {
	spotMap := make(map[int]*ParkingSpot)

	for spotId := 1; spotId <= spotCount[vehicleType]; spotId ++ {
		spotMap[spotId] = NewParkingSpot(spotId, vehicleType)
	}

	return spotMap
}

func (pl *ParkLevel) FindParkingSpot(vehicleType vehicles.VehicleType) *ParkingSpot {
	for _, spot := range pl.ParkingSpots[vehicleType] {
		if spot.IsAvailable() {
			return spot
		}
	}
	return nil
}

func (pl *ParkLevel) DisplayAvailability() {
	fmt.Println("Floor ID: ", pl.FloorId)
	for vType, spotMap := range pl.ParkingSpots {
		cnt := 0
		for _, spot := range spotMap {
			if spot.IsAvailable() {
				cnt++
			}
		}
		fmt.Println(vehicles.VehicleTypeMap[vType], " Spots available : ", cnt)
	}
}