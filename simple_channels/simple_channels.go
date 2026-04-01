package simple_channels

import (
	"fmt"
	"strings"
)

// Designating receive and send only channels
func shout(ping <-chan string, pong chan<- string) {
	for {

		// Read from ping channel. Note that the GoRoutine waits here -- it blocks till something is recieved on this channel
		s, ok := <-ping

		if !ok {
			// Channel is closed
			return
		}
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func ChannelDemo() {
	ping := make(chan string)

	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit):")

	for {
		// Print prompt
		fmt.Print("> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput

		// Wait for response
		response := <-pong
		fmt.Println("Response:", response)
	}

	fmt.Println("Goodbye!")
	close(ping)
	close(pong)
}
