package main

import "go.uber.org/fx"

func main() {
	fx.New(
		fx.Provide(
			NewConfig,
			NewTaskRepository,
			NewTaskService,
			NewTaskHandler,
			NewRouter,
			NewHTTPServer,
		),
		fx.Invoke(StartServer),
	).Run()
}
