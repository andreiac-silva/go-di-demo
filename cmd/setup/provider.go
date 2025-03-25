package setup

import (
	"github.com/andreiac-silva/go-di-demo/cmd/setup/database"
	"github.com/andreiac-silva/go-di-demo/cmd/setup/server"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	database.NewPgConn,
	server.NewHTTPServer,
	wire.Value("tenant"),
)
