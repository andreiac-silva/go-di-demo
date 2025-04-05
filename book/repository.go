package book

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

func (r repository) Save(ctx context.Context, book domain.Book) (int64, error) {
	var bookID int64

	err := r.db.QueryRow(ctx, `
		INSERT INTO books (isbn, title, author, genre, published_at) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id
		`, book.ISBN, book.Title, book.Author, book.Genre, book.PublishedAt).Scan(&bookID)

	return bookID, err
}

func (r repository) Get(ctx context.Context, id int64) (domain.Book, error) {
	var book domain.Book

	err := r.db.QueryRow(ctx, `
		SELECT b.id, b.isbn, b.title, b.author, b.genre, b.published_at, b.created_at, b.updated_at
		FROM books b 
		WHERE id = $1 
		AND b.deleted_at IS NULL
	`, id).Scan(
		&book.ID,
		&book.ISBN,
		&book.Title,
		&book.Author,
		&book.Genre,
		&book.PublishedAt,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	return book, err
}
