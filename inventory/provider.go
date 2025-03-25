package inventory

import (
	"github.com/andreiac-silva/go-di-demo/domain"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewRepository,
	NewService,
	wire.Bind(new(domain.InventoryRepository), new(*repository)),
	wire.Bind(new(domain.InventoryService), new(*service)),
)
