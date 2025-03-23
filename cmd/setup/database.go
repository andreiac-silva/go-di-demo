package setup

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func NewPgConn(ctx context.Context) *pgx.Conn {
	db, err := pgx.Connect(ctx, PostgresDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return db
}
