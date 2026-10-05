package main

import (
	"fmt"
)

type Bakery struct {
	ovenTemperature int
	ingredients     []string
}

func (b *Bakery) Bake() {
	switch {
	case b.ovenTemperature < 350:
		fmt.Println("Oven temperature is too low for baking.")
	case b.ovenTemperature > 450:
		fmt.Println("Oven temperature is too high for baking.")
	default:
		fmt.Printf("Baking at %d degrees with ingredients: %v\n", b.ovenTemperature, b.ingredients)
	}
	for _, ingredient := range b.ingredients {
		switch ingredient {
		case "flour":
			fmt.Println("Adding flour.")
		case "sugar":
			fmt.Println("Adding sugar.")
		case "eggs":
			fmt.Println("Adding eggs.")
		default:
			fmt.Printf("Adding unknown ingredient: %s\n", ingredient)
		}

	}
	fmt.Println("Baking complete!")
}

func main() {
	fmt.Println("This is the old method.")
	bakery := Bakery{
		ovenTemperature: 375,
		ingredients:     []string{"flour"},
	}
	bakery.Bake()
	bakery2 := Bakery{
		ovenTemperature: 300,
		ingredients:     []string{"flour", "coconut", "eggs"},
	}
	bakery2.Bake()
}
