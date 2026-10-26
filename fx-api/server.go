package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"go.uber.org/fx"
)

func NewHTTPServer(config Config, mux *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: mux,
	}
}

func StartServer(lifecycle fx.Lifecycle, server *http.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Printf("HTTP server stopped: %v", err)
				}
			}()
			log.Printf("HTTP server listening on %s", server.Addr)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
