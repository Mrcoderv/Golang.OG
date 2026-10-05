package main



type Engine struct{}  // this is a dependency of the Car struct

func (e Engine) Start() {
	println("Engine started")
}