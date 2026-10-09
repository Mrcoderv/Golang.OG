package main

import (
	"go.uber.org/fx"
)

func Run(service *WriteService) error {
	return service.Execute()
}

func main() {
	app := fx.New(
		fx.Provide(
			fx.Annotate(
				NewFileWriter,
				fx.As(new(Writer)),
			),
			NewWriteService,
		),
		fx.Invoke(Run),
	)

	app.Run()
}
