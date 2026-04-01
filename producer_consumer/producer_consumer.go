package producer_consumer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Recieved order #%d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}

		total++

		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizzaNumber, delay)

		//delay for a while
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("We ran out of ingredients for pizza #%d", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("We were attacked by a moose while making pizza #%d", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza #%d is ready", pizzaNumber)
		}

		p := PizzaOrder{pizzaNumber, msg, success}

		return &p

	}

	return &PizzaOrder{pizzaNumber, "No more pizzas", false}
}

func makePizzaria(pizzaMaker *Producer) {
	//keep track of which pizza we are making
	var i = 0

	//run forever until we are told to stop
	//try to make pizzas

	for {
		currentPizza := makePizza(i)

		if currentPizza != nil {
			i = currentPizza.pizzaNumber

			select {
			// We tried to make a pizza (we are sending something to the pizzaMaker.data channel)
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func ProducerConsumer() {
	// initiate and seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// print out a message
	color.Cyan("The Pizzaria is open for business!")
	color.Cyan("----------------------------------")

	//create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//run producer in background in its own go routine
	go makePizzaria(pizzaJob)

	//crete and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery\n", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Done making pizzas!")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing producer: %s", err)
			}
		}
	}
	color.Cyan("-----------------")
	color.Cyan("Done for the day!")

	color.Cyan("Made %d pizzas, %d failed, %d total", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("It was a disaster!")
	case pizzasFailed >= 6:
		color.Red("It was not a great day...")
	case pizzasFailed >= 4:
		color.Yellow("It was an okay day...")
	case pizzasFailed >= 2:
		color.Yellow("It was a pretty good day...")
	default:
		color.Green("It was a great day!")
	}
}
