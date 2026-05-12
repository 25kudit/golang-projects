package kitchen

import (
	"log"
	"pizzeria/models"
	"strconv"
	"time"
)

var (
	prepQueue chan models.Order
	ovenQueue chan models.Order
	deliveryQueue chan models.Order
)

func Init() {
	prepQueue = make(chan models.Order, 10)
	ovenQueue = make(chan models.Order, 10)
	deliveryQueue = make(chan models.Order, 10)

	for i := range 3 {
		go startChef("C-" + strconv.Itoa(i))
	}

	for i := range 2 {
		go startOven("O-" + strconv.Itoa(i))
	}

	for i := range 2 {
		go deliverOrder("D-" + strconv.Itoa(i))
	}
}

func PlaceOrder(order models.Order) {
	prepQueue <- order
}

func startChef(chefId string) {
	for order := range prepQueue {
		log.Println("started preparing order id:", order.Id, " by chef id:", chefId)
		time.Sleep(5 * time.Second)
		order.Status = "prepped"
		log.Println("order prepped, id:", order.Id, " by chef id:", chefId)
		ovenQueue <- order
	}
}

func startOven(ovenId string) {
	for order := range ovenQueue {
		log.Println("started baking order id:", order.Id, " in oven id:", ovenId)
		time.Sleep(7 * time.Second)
		order.Status = "done"
		log.Println("order ready, id:", order.Id, " in oven id:", ovenId)
		deliveryQueue <- order
	}
}

func deliverOrder(driverId string) {
	for order := range deliveryQueue {
		log.Println("Delivery partner assigned to order id:", order.Id, " driver id:", driverId)
		time.Sleep(10 * time.Second)
		order.Status = "delivered"
		log.Println("order delivered, id:", order.Id, " by driver id:", driverId)
	}
}