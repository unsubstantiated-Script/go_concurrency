package sleeping_barber

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// variables

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func SleepingBarber() {
	// seed rando number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))
	//print welcome
	color.Yellow("Welcome to the Sleeping Barber Shop!")
	color.Yellow("------------------------------------")

	// create channels
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create data struct for barber shop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The shop is now open!")

	// add barbers to shop
	shop.addBarber("Frank Zappa")
	shop.addBarber("Chas Swink")
	shop.addBarber("Gandalf Greymane")
	shop.addBarber("Zoot Alour")
	shop.addBarber("Chili Rib")

	// start barbershop as a goroutine

	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closedShopForDay()
		closed <- true
	}()

	// add clients
	i := 1

	go func() {
		for {
			randomMillisecs := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMillisecs)):
				shop.addClient(fmt.Sprintf("Client %d", i))
				i++
			}
		}
	}()
	// block until the barbershop is closed
	<-closed

}
