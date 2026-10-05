package main

func main() {

	engine := Engine{}

	car := Car{
		engine: engine,
	}
	car.Start()
}
