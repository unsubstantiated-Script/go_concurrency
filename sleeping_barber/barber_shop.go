package sleeping_barber

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	ClientsChan     chan string
	BarbersDoneChan chan bool
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients", barber)

		for {
			//if there are no clients, barber sleeps
			if len(shop.ClientsChan) == 0 {
				color.Yellow("%s is sleeping", barber)
				isSleeping = true
			}

			client, shopOpenCheck := <-shop.ClientsChan

			if shopOpenCheck {
				if isSleeping {
					color.Yellow("%s wakes %s up", client, barber)
					isSleeping = false
				}
				// cut hair
				shop.cutHair(barber, client)
			} else {
				// shop is closed, so send the barber home and close this goroutine
				shop.sendBarberHome(barber)
				return
			}

		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s cuts hair for %s", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is done cutting hair for %s", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is leaving the shop", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closedShopForDay() {
	color.Cyan("The shop is closed for the day")

	close(shop.ClientsChan)
	shop.Open = false

	for i := 1; i <= shop.NumberOfBarbers; i++ {
		<-shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)
	color.Green("The shop is now closed")
	color.Green("----------------------")
}

func (shop *BarberShop) addClient(client string) {
	color.Green("Client %s arrived at the shop", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("Client %s was added to the queue", client)
		default:
			color.Red("The shop is full, so %s leaves", client)
		}
	} else {
		color.Red("The shop is closed, so %s leaves", client)
	}
}
