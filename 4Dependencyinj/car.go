package main

import "fmt"

type Car struct {  
	engine Engine
}

func (c Car) Start() {
	c.engine.Start()
	fmt.Println("Car started")
}
