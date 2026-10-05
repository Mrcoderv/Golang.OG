type Payment interface { // pay is the interface that defines the payment method
	Pay(amount float64) error
}

type Esewa struct{} // esewa is the struct that implements the Payment interface

func (e Esewa) Pay(amount float64) error {
	fmt.Println("Paid with eSewa")
	return nil
}

type Khalti struct{} // another struct that implements the Payment interface

func (k Khalti) Pay(amount float64) error {
	fmt.Println("Paid with Khalti")
	return nil
}

// type Stripe struct{}  // a struct that implements the Payment interface

// func (s Stripe) Pay(amount float64) error {
//     fmt.Println("Paid with Stripe")
//     return nil
// }