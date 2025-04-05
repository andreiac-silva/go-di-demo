package book

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/andreiac-silva/go-di-demo/domain"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service domain.BookService
}

func NewHandler(service domain.BookService) *Handler {
	return &Handler{service: service}
}

func (h Handler) Routes(router *gin.Engine) {
	router.POST("/books", h.create)
	router.GET("/books/:id", h.get)
}

func (h Handler) create(c *gin.Context) {
	book := domain.Book{}
	if err := c.BindJSON(&book); err != nil {
		_ = c.Error(fmt.Errorf("failed to parse book: %w", err))
		return
	}

	id, err := h.service.Save(c, book)
	if err != nil {
		_ = c.Error(fmt.Errorf("failed to save book: %w", err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h Handler) get(c *gin.Context) {
	idStr := c.Params.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		_ = c.Error(fmt.Errorf("failed parse book id: %w", err))
		return
	}

	book, err := h.service.Get(c, id)
	if err != nil {
		_ = c.Error(fmt.Errorf("failed to get book: %w", err))
		return
	}

	c.JSON(http.StatusOK, book)
}
