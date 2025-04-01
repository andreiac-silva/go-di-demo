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
	"github.com/andreiac-silva/go-di-demo/cmd/setup/database"
	"github.com/andreiac-silva/go-di-demo/cmd/setup/server"
	"github.com/andreiac-silva/go-di-demo/inventory"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := database.NewPgConn(ctx, "tenant")
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	defer db.Close(ctx)

	inventoryRepo := inventory.NewRepository(db)
	inventoryService := inventory.NewService(inventoryRepo)
	bookRepo := book.NewRepository(db)
	bookService := book.NewService(bookRepo, inventoryService)
	handler := book.NewHandler(bookService)
	srv := server.NewHTTPServer(handler)

	log.Printf("starting server at '%s'", srv.Addr)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error on starting server: %v", err)
		}
	}()

	trapSignalShutdown(func() {
		log.Println("shutting down server...")
		if err := srv.Shutdown(ctx); err != nil {
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
