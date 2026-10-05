package main

import "fmt"

// this is the depandency
type Oven interface {
	temperature() string
}
type Ingredients interface {
	mix() []string
}

type LowTemperatureOven struct{}

func (o LowTemperatureOven) temperature() string {
	return "Oven temperature is too low for baking."
}

type NormalTemperatureOven struct{}

func (o NormalTemperatureOven) temperature() string {
	return "Baking at normal temperature."
}

type HighTemperatureOven struct{}

func (o HighTemperatureOven) temperature() string {
	return "Oven temperature is too high for baking."
}

type Flour struct{}

func (f Flour) mix() []string {
	return []string{"Adding flour."}
}

type Coconut struct{}

func (c Coconut) mix() []string {
	return []string{"Adding coconut."}
}

type Eggs struct{}

func (e Eggs) mix() []string {
	return []string{"Adding eggs."}
}

type Bakery struct {
	oven        Oven
	ingredients Ingredients
}

func (b Bakery) Bake() {
	fmt.Println(b.oven.temperature())

	for _, ingredient := range b.ingredients.mix() {
		fmt.Println(ingredient)
	}

	fmt.Println("Baking...")
	fmt.Println()
}

// dependency part close.

func main() {
	fmt.Println("Using DI.")

	oven := NormalTemperatureOven{}

	ingredients := Flour{}

	bakery := Bakery{ //dependency injection
		oven:        oven,
		ingredients: ingredients,
	}

	bakery.Bake()

	oven2 := LowTemperatureOven{}
	ingredients2 := Coconut{}

	bakery2 := Bakery{
		oven:        oven2,
		ingredients: ingredients2,
	}

	bakery2.Bake()
}
