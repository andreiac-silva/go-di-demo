package book

import (
	"github.com/andreiac-silva/go-di-demo/domain"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewRepository,
	NewService,
	NewHandler,
	wire.Bind(new(domain.BookRepository), new(*repository)),
	wire.Bind(new(domain.BookService), new(*service)),
)
