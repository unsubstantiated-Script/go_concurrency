package second_example

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func SecondExample() {
	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("Hello, universe!")
	go updateMessage("Hello, cosmos?")
	wg.Wait()

	fmt.Println(msg)
}

//var msg string
//var wg sync.WaitGroup
//
//func updateMessage(s string, m *sync.Mutex) {
//	defer wg.Done()
//
//	m.Lock()
//	msg = s
//	m.Unlock()
//}
//
//func SecondExample() {
//	msg = "Hello, world!"
//
//	var mutex sync.Mutex
//
//	wg.Add(2)
//	go updateMessage("Hello, universe!", &mutex)
//	go updateMessage("Hello, cosmos?", &mutex)
//	wg.Wait()
//
//	fmt.Println(msg)
//}

func AnotherSecondExample() {
	var bankBalance int
	var balance sync.Mutex

	fmt.Printf("Your bank balance is %d\n", bankBalance)
	fmt.Println()

	incomes := []Income{
		{Source: "Salary", Amount: 10000},
		{Source: "Investment", Amount: 50000},
		{Source: "Charity", Amount: 2000},
		{Source: "Cherries", Amount: 10000},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("Week %d: You earned %d.00 from %s\n", week, income.Amount, income.Source)
			}

		}(i, income)
	}

	wg.Wait()
	fmt.Printf("Your bank balance is %d.00\n", bankBalance)
	fmt.Println()
}
