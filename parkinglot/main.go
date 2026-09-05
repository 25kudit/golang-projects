package main

import (
	"fmt"
	"parkinglot/vehicles"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {
	parkingLot := GetParkingLotInstance()

	parkingLot.AddParkLevel(0)
	parkingLot.AddParkLevel(1)
	parkingLot.DisplayAvailability()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go parkCar(i, parkingLot)
	}

	wg.Wait()

	parkingLot.DisplayAvailability()

	truck := vehicles.NewTruck("truck-1")
	truckTicket, _ := parkingLot.ParkVehicle(truck)

	time.Sleep(5 * time.Second)
	
	parkingLot.DisplayAvailability()

	charges := parkingLot.UnparkVehicle(truckTicket)

	fmt.Printf("total charges for %s: %.2f", truck.LicensePlate, charges)

}

func parkCar(i int, parkingLot *ParkingLot) {
	defer wg.Done()

	car := vehicles.NewCar(fmt.Sprintf("car-%d", i))
	ticket, err := parkingLot.ParkVehicle(car)
	if err != nil {
		fmt.Printf("Failed to park %s: %v\n", car.LicensePlate, err)
		return
	}

	fmt.Printf("%s car parked successfully, Ticket entry time: %s\n", car.LicensePlate, ticket.EntryTime)
}