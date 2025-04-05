package setup

import (
	"context"
	"time"

	"github.com/andreiac-silva/go-di-demo/cmd/setup/database"
	"github.com/andreiac-silva/go-di-demo/cmd/setup/server"

	"github.com/google/wire"
	"go.uber.org/dig"
	"go.uber.org/fx"
)

// Provider is a wire provider set to initiate database and http server.
var Provider = wire.NewSet(
	database.NewPgConn,
	server.NewRouterContainer,
	server.NewHTTPServerForWire,
	wire.Value("tenant"),
)

// Module is a fx module set to initiate database and http server.
var Module = fx.Module("app",
	fx.Provide(
		func(lc fx.Lifecycle) context.Context {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			lc.Append(fx.Hook{
				OnStop: func(context.Context) error {
					cancel()
					return nil
				},
			})
			return ctx
		},
		func() string { return "tenant" },
		database.NewPgConn,
	),
)

// Provide is a function that provides the database and http server to the dig container.
func Provide(container *dig.Container) error {
	if err := container.Provide(func() string { return "tenant" }); err != nil {
		return err
	}
	return container.Provide(database.NewPgConn)
}
