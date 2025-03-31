package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"go.uber.org/fx"
)

func RegisterHTTPServerForFx(lc fx.Lifecycle, srv *http.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Printf("starting server at '%s'", srv.Addr)
			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("error on starting server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("shutting down server...")
			return srv.Shutdown(ctx)
		},
	})
}
