package dining_philosophers

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// DiningPhilosophers. Five guys sitting in a circle at a table. Five forks. To eat, each needs two forks. Neighbors can't share.

// Philosopher struct for a philosopher and the forks they need.
type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

// lit of all the philosophers.
var philosophers = []Philosopher{
	{"Albert", 4, 0},
	{"Bob", 0, 1},
	{"Carl", 1, 2},
	{"Douglas", 2, 3},
	{"Erik", 3, 4},
}

var hunger = 3 // how times does a philosopher eat?
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

var orderMutex sync.Mutex
var orderFinished []string

func DiningPhilosophers() {
	// print out a welcome message
	fmt.Println("Welcome to the Dining Philosophers Problem!")
	fmt.Println("-------------------------------------------")
	fmt.Println("The table is ready!")

	time.Sleep(sleepTime)

	// start the meal
	dine()

	// print out finished message
	fmt.Println("The table is empty.")
	fmt.Println("-------------------------------------------")
	fmt.Println("The meal is finished!")

	// *** added this
	time.Sleep(sleepTime)
	fmt.Printf("Order finished: %s.\n", strings.Join(orderFinished, ", "))
}

func dine() {

	// This syncs up when all the philosophers are done eating.
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// This syncs up when all the philosophers are seated.
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	//forks is a map of all five forks.
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s is seated at the table.\n", philosopher.name)
	seated.Done()

	seated.Wait()

	for i := hunger; i > 0; i-- {

		// get a lock on both forks
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
		}

		fmt.Printf("\t%s has both forks and is chawing down.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is now ruminating.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		fmt.Printf("\t%s puts down the left fork.\n", philosopher.name)
		forks[philosopher.rightFork].Unlock()
		fmt.Printf("\t%s puts down the right fork.\n", philosopher.name)
	}

	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()

	fmt.Printf("%s is done eating.\n", philosopher.name)
	fmt.Printf("%s left the table.\n", philosopher.name)
}
