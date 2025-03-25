package database

import (
	"context"
	"log"

	"github.com/andreiac-silva/go-di-demo/cmd/setup/env"

	"github.com/jackc/pgx/v5"
)

func NewPgConn(ctx context.Context, tenant string) (*pgx.Conn, error) {
	log.Printf("it is only to simulate wire.Value. parameter: '%s'", tenant)
	return pgx.Connect(ctx, env.PostgresDSN)
}
