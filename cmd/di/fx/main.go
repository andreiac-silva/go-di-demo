package main

import (
	"github.com/andreiac-silva/go-di-demo/api"
	"github.com/andreiac-silva/go-di-demo/book"
	"github.com/andreiac-silva/go-di-demo/cmd/setup"
	"github.com/andreiac-silva/go-di-demo/cmd/setup/server"
	"github.com/andreiac-silva/go-di-demo/inventory"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.NopLogger,

		inventory.Module,
		book.Module,
		setup.Module,

		fx.Provide(
			fx.Annotate(book.NewHandler, fx.As(new(api.Router)),
				fx.ResultTags(`group:"routers"`)),
		),

		fx.Provide(
			fx.Annotate(
				server.NewHTTPServer,
				fx.ParamTags(`group:"routers"`),
			),
		),

		fx.Invoke(server.RegisterHTTPServerForFx),
	)

	app.Run()
}
