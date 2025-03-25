package inventory

import (
	"context"

	"github.com/andreiac-silva/go-di-demo/domain"

	"github.com/jackc/pgx/v5"
)

type repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) *repository {
	return &repository{db: db}
}

func (r repository) Save(ctx context.Context, inventory domain.Inventory) (int64, error) {
	var inventoryID int64

	err := r.db.QueryRow(ctx, `
		INSERT INTO inventories (book_id, quantity) 
		VALUES ($1, $2) 
		RETURNING id
		`, inventory.BookID, inventory.Quantity).Scan(&inventoryID)

	return inventoryID, err
}

func (r repository) Get(ctx context.Context, id int64) (domain.Inventory, error) {
	var inventory domain.Inventory

	err := r.db.QueryRow(ctx, `
		SELECT i.id, i.quantity, i.created_at, i.updated_at
		FROM inventories i 
		WHERE id = $1 
		AND i.deleted_at IS NULL
	`, id).Scan(
		&inventory.ID,
		&inventory.Quantity,
		&inventory.CreatedAt,
		&inventory.UpdatedAt,
	)

	return inventory, err
}
