package book

import (
	"context"
	"fmt"

	"github.com/andreiac-silva/go-di-demo/domain"
)

type service struct {
	repository       domain.BookRepository
	inventoryService domain.InventoryService
}

func NewService(repository domain.BookRepository, inventoryService domain.InventoryService) *service {
	return &service{
		repository:       repository,
		inventoryService: inventoryService,
	}
}

func (s service) Save(ctx context.Context, book domain.Book) (int64, error) {
	// Whoops! No transaction wrapping here - bold move! 😆
	// But hey, let's not get sidetracked. The real star of the show is dependency injection!
	id, err := s.repository.Save(ctx, book)
	if err != nil {
		return 0, fmt.Errorf("failed to save book: %w", err)
	}

	inventory := domain.NewInventory(id, book.Inventory.Quantity)
	if _, err = s.inventoryService.Save(ctx, inventory); err != nil {
		return 0, fmt.Errorf("failed to create inventory for book ID %d: %w", id, err)
	}

	return id, nil
}

func (s service) Get(ctx context.Context, id int64) (domain.Book, error) {
	book, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Book{}, fmt.Errorf("failed to get book: %w", err)
	}

	book.Inventory, err = s.inventoryService.Get(ctx, id)
	if err != nil {
		return domain.Book{}, fmt.Errorf("failed to get inventory for book ID %d: %w", id, err)
	}

	return book, nil
}
