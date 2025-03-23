package domain

import (
	"context"
	"time"
)

type Book struct {
	ID          int64     `json:"id"`
	ISBN        string    `json:"isbn"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Genre       string    `json:"genre"`
	Inventory   Inventory `json:"inventory"`
	PublishedAt time.Time `json:"published_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type BookRepository interface {
	Save(ctx context.Context, book Book) (int64, error)
	Get(ctx context.Context, id int64) (Book, error)
}

type BookService interface {
	Save(ctx context.Context, book Book) (int64, error)
	Get(ctx context.Context, id int64) (Book, error)
}
