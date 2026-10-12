-- Active: 1790580078615@@127.0.0.1@5432@taskmanager
package main

import "fmt"

// Interface
type Payment interface { // defines behaviour
	Pay(amount int)
}

// implementation 1
type Esewa struct{} // the autal work..

func (e Esewa) Pay(amount int) {
	fmt.Println("Paid using eSewa:", amount)
}

// Implementation 2
type Khalti struct{}

func (k Khalti) Pay(amount int) { //
	fmt.Println("Paid using Khalti:", amount)
}

// Function that takes the interface as a parameter
func ProcessPayment(p Payment) {
	p.Pay(1000)
}

type PaymentService struct {
	payment Payment
}

// “
func NewPaymentService(p Payment) *PaymentService { // dependency injection suing the constructor
	return &PaymentService{
		payment: p,
	} // NEW KEY WORD IS USED TO MEKE THE CONSTRUCTOR FUNCTION THAT RETURNS A POINTER TO THE STRUCTs
}

func (s *PaymentService) MakePayment(amount int) { // this is the implementation that depends on the interface
	s.payment.Pay(amount)
}

func main() {
	esewa := Esewa{}
	khalti := Khalti{}

	ProcessPayment(esewa) // using the interface to call the implementation
	ProcessPayment(khalti)

	// using the dependency injection to call the implementation
	task := NewPaymentService(esewa)  // here we manually inject the dependency by passing the implementation to the constructor function
	task.MakePayment(1880)
	NewPaymentService(khalti).MakePayment(1000)
}
