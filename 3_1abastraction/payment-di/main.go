package main

import (
	"fmt"

	"go.uber.org/fx"
)

type Payment interface {
	Pay(amount int)
}

type Esewa struct{}

func (e Esewa) Pay(amount int) {
	fmt.Println("Paid using eSewa:", amount)
}

type Khalti struct{}

func (k Khalti) Pay(amount int) {
	fmt.Println("Paid using Khalti:", amount)
}

// Output for the eSewa provider.
type EsewaResult struct {
	fx.Out
	Payment Payment `group:"payments"`
}

func NewEsewa() EsewaResult {
	return EsewaResult{
		Payment: Esewa{},
	}
}

// Output for the Khalti provider.
type KhaltiResult struct {
	fx.Out
	Payment Payment `group:"payments"`
}

func NewKhalti() KhaltiResult {
	return KhaltiResult{
		Payment: Khalti{},
	}
}

// Collect all implementations from the payments group.
type PaymentParams struct {
	fx.In
	Payments []Payment `group:"payments"`
}

func ProcessPayments(p PaymentParams) {
	for _, payment := range p.Payments {
		payment.Pay(10330)
	}
}

func main() {
	fx.New(
		fx.Provide(
			NewEsewa,
			NewKhalti,
		),
		fx.Invoke(ProcessPayments),
	).Run()
}
