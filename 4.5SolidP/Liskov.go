// liskov substitution principle
// means Replace implementation → it should still work as expected.
package main

import "fmt"

type Payment interface { // pay is the interface that defines the payment method
	Pay(amount float64) error
}

// // old  implementation
// func Payment(p Payment, amount float32) {
// 	p.Pay(amount)
// }

// replacing  the old with new  implementation
func processPayment(p Payment, amount float64) {
	p.Pay(amount)
}

type Esewa struct{} // esewa is the struct that implements the Payment interface

func (e Esewa) Pay(amount float64) error {
	fmt.Println("Paid by eSewa")
	return nil
}

type Khalti struct{} // another struct that implements the Payment interface

func (k Khalti) Pay(amount float64) error {
	fmt.Println("Paid by Khalti")
	return nil
}
func main() {
	esewa := Esewa{}
	khalti := Khalti{}

	processPayment(esewa, 150.0)  // Paid with eSewa
	processPayment(khalti, 250.0) // Paid with Khalti
}
