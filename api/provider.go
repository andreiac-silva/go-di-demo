package api

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewRouterContainer,
)
