package main

import (
	"fmt"

	"go.uber.org/fx"
)

// Interface
type Payment interface {
	Pay(amount int)
}

// implementation 1
type Esewa struct{}

func (e Esewa) Pay(amount int) {
	fmt.Println("Paid using eSewa:", amount)
}

// Implementation 2
type Khalti struct{}

func (k Khalti) Pay(amount int) { //
	fmt.Println("Paid using Khalti:", amount)
}

// Function that takes the interface as a parameter
func ProcessPayment(p Payment) { // abstraction..
	p.Pay(1000)
}

func main() {
	fx.New(
		fx.Provide(
			func() Payment {
				return Esewa{}
			},
			func() Payment {
				return Khalti{}

			},
		),
		fx.Invoke(
			func(p Payment) {
				ProcessPayment(p)
			},
		),
	).Run()
	esewa := Esewa{}
	khalti := Khalti{}

	ProcessPayment(esewa)
	ProcessPayment(khalti)
}
