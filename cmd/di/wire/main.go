package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server, err := InitApplication(ctx)
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}

	log.Printf("starting server at '%s'", server.Addr)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error on starting server: %v", err)
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
