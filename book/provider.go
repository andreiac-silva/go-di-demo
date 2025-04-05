package book

import (
	"github.com/andreiac-silva/go-di-demo/api"
	"github.com/andreiac-silva/go-di-demo/domain"

	"github.com/google/wire"
	"go.uber.org/dig"
	"go.uber.org/fx"
)

// Provider is the wire provider set for the book module.
var Provider = wire.NewSet(
	NewRepository,
	NewService,
	NewHandler,
	wire.Bind(new(domain.BookRepository), new(*repository)),
	wire.Bind(new(domain.BookService), new(*service)),
)

// Module is the fx module for the book module.
var Module = fx.Module("book",
	fx.Provide(
		fx.Annotate(NewRepository, fx.As(new(domain.BookRepository))),
		fx.Annotate(NewService, fx.As(new(domain.BookService))),
		fx.Annotate(NewHandler, fx.As(new(api.Router)),
			fx.ResultTags(`group:"routers"`)),
	),
)

// Provide is a function that set up the dig book dependency container.
func Provide(container *dig.Container) error {
	if err := container.Provide(NewRepository, dig.As(new(domain.BookRepository))); err != nil {
		return err
	}
	if err := container.Provide(NewService, dig.As(new(domain.BookService))); err != nil {
		return err
	}
	return container.Provide(NewHandler, dig.As(new(api.Router)), dig.Group("routers"))
}
