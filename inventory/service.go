package inventory

import (
	"context"

	"github.com/andreiac-silva/go-di-demo/domain"
)

type service struct {
	repository domain.InventoryRepository
}

func (s service) Save(ctx context.Context, inventory domain.Inventory) (int64, error) {
	return s.repository.Save(ctx, inventory)
}

func (s service) Get(ctx context.Context, id int64) (domain.Inventory, error) {
	return s.repository.Get(ctx, id)
}
