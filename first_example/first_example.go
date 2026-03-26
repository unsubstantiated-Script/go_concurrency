package first_example

import (
	"fmt"
	"sync"
)

func RunWaitGroup() {
	var wg sync.WaitGroup

	words := []string{"Hello", "World", "Derp", "Bilbo", "Gandalf"}

	wg.Add(len(words))

	for i, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, word), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomething("Hello, World!", &wg)
}

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
