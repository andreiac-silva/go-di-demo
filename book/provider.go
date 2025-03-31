package book

import (
	"github.com/andreiac-silva/go-di-demo/domain"

	"github.com/google/wire"
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
	),
)
