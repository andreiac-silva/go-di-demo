//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"net/http"

	"github.com/andreiac-silva/go-di-demo/api"
	"github.com/andreiac-silva/go-di-demo/book"
	"github.com/andreiac-silva/go-di-demo/cmd/setup"
	"github.com/andreiac-silva/go-di-demo/inventory"

	"github.com/google/wire"
)

var Container = wire.NewSet(
	setup.Provider,
	book.Provider,
	inventory.Provider,
	api.Provider,
)

func InitApplication(ctx context.Context) (*http.Server, error) {
	wire.Build(Container)
	return nil, nil
}
