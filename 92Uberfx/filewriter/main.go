package main

import (
	"go.uber.org/fx"
)

func Run(service *WriteService) {
	if err := service.Execute(); err != nil {
		panic(err)
	}
}

func main() {
	fx.New(
		fx.Provide(
			NewFileWriter,
			NewWriteService,
		),
		fx.Invoke(Run),
	).Run()
}
