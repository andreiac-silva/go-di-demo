package domain

import (
	"context"
	"time"
)

type Inventory struct {
	ID        int64     `json:"id"`
	BookID    int64     `json:"book_id"`
	Quantity  int       `json:"quantity"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewInventory(bookID int64, quantity int) Inventory {
	return Inventory{
		BookID:   bookID,
		Quantity: quantity,
	}
}

type InventoryRepository interface {
	Save(ctx context.Context, inventory Inventory) (int64, error)
	Get(ctx context.Context, id int64) (Inventory, error)
}

type InventoryService interface {
	Save(ctx context.Context, inventory Inventory) (int64, error)
	Get(ctx context.Context, id int64) (Inventory, error)
}
