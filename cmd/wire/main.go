package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	server, err := InitApplication(ctx)
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}

	log.Printf("starting server on '%s'", server.Addr)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	trapSignalShutdown(func() {
		log.Println("shutting down server...")
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("server shutdown error: %v", err)
		}
		log.Println("server stopped")
	})
}

func trapSignalShutdown(onShutdown func()) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	onShutdown()
}
