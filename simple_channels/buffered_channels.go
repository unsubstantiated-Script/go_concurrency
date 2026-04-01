package simple_channels

import (
	"fmt"
	"time"
)

func listenToChannel(ch chan int) {
	for {
		// print a got data message
		i := <-ch
		fmt.Println("Got data from channel: ", i)

		// simulate some work
		time.Sleep(1 * time.Second)
	}
}

func BufferedChannelDemo() {

	// This is an example of a buffered channel. It will only accept 10 items before it blocks. But this can speed up the process of sending data.
	// Buffers are useful when you want to make sure that you are not overloading the system with too many messages.
	// Buffers are useful when you know how many go routines you will be running.
	// Limit the number of messages that can be sent to the channel.
	ch := make(chan int, 10)

	go listenToChannel(ch)

	for i := 0; i < 100; i++ {
		fmt.Println("Sending data to channel: ", i)
		ch <- i
		fmt.Println("sent", i, "to channel")
	}

	fmt.Println("Done")
	close(ch)
}
