package inventory

import (
	"github.com/andreiac-silva/go-di-demo/domain"

	"github.com/google/wire"
	"go.uber.org/dig"
	"go.uber.org/fx"
)

// Provider is a wire provider set for the inventory domain.
var Provider = wire.NewSet(
	NewRepository,
	NewService,
	wire.Bind(new(domain.InventoryRepository), new(*repository)),
	wire.Bind(new(domain.InventoryService), new(*service)),
)

// Module is a fx module for the inventory domain.
var Module = fx.Module("inventory",
	fx.Provide(
		fx.Annotate(NewRepository, fx.As(new(domain.InventoryRepository))),
		fx.Annotate(NewService, fx.As(new(domain.InventoryService))),
	),
)

// Provide is a function that provides the inventory module to the dig container.
func Provide(container *dig.Container) error {
	if err := container.Provide(NewRepository, dig.As(new(domain.InventoryRepository))); err != nil {
		return err
	}
	return container.Provide(NewService, dig.As(new(domain.InventoryService)))
}
