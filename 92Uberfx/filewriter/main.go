package main

import (
	"context"

	"go.uber.org/fx"
)

func Run(service *WriteService) {
	if err := service.Execute(); err != nil {
		panic(err)
	}
}

func main() {
	app := fx.New(
		fx.Provide(
			NewFileWriter,
			NewWriteService,
		),
		fx.Invoke(Run), // creating the uberfx app and registering the depencency.
	)

	if err := app.Start(context.Background()); err != nil {
		panic(err)
	}

	defer app.Stop(context.Background())
}
