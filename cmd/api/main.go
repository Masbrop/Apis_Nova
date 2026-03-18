package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"apis_nova/internal/infrastructure/bootstrap"
)

func main() {
	app, err := bootstrap.NewApplication()
	if err != nil {
		log.Fatalf("bootstrap application: %v", err)
	}
	defer func() {
		if closeErr := app.Close(); closeErr != nil {
			log.Printf("close resources: %v", closeErr)
		}
	}()

	serverErr := make(chan error, 1)

	go func() {
		serverErr <- app.Server.ListenAndServe()
	}()

	log.Printf("api listening on %s", app.Config.HTTPAddress())

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case err := <-serverErr:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http server: %v", err)
		}
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), app.Config.ShutdownTimeout)
		defer cancel()

		if err := app.Server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("shutdown server: %v", err)
		}
	}
}
