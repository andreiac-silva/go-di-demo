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

	"github.com/andreiac-silva/go-di-demo/book"
	"github.com/andreiac-silva/go-di-demo/cmd/setup"
	"github.com/andreiac-silva/go-di-demo/cmd/setup/server"
	"github.com/andreiac-silva/go-di-demo/inventory"

	"go.uber.org/dig"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	container := dig.New()
	if err := container.Provide(func() context.Context { return ctx }); err != nil {
		log.Fatalf("failed to inform context: %v", err)
	}
	if err := setup.Provide(container); err != nil {
		log.Fatalf("failed to init base dependencies: %v", err)
	}
	if err := book.Provide(container); err != nil {
		log.Fatalf("failed to init book dependencies: %v", err)
	}
	if err := inventory.Provide(container); err != nil {
		log.Fatalf("failed to init inventory dependencies: %v", err)
	}
	if err := container.Provide(server.NewHTTPServerForDig); err != nil {
		log.Fatalf("failed to init server: %v", err)
	}

	err := container.Invoke(func(s *http.Server) {
		log.Printf("starting server at '%s'", s.Addr)

		go func() {
			if err := s.ListenAndServe(); err != nil && !errors.As(err, &http.ErrServerClosed) {
				log.Fatalf("error on starting server: %v", err)
			}
		}()

		trapSignalShutdown(func() {
			log.Println("shutting down server...")
			if err := s.Shutdown(ctx); err != nil {
				log.Fatalf("server shutdown error: %v", err)
			}
			log.Println("server stopped")
		})
	})

	if err != nil {
		log.Fatalf("error on starting server: %v", err)
	}
}

func trapSignalShutdown(onShutdown func()) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	onShutdown()
}
