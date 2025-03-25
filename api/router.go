package api

import (
	"github.com/andreiac-silva/go-di-demo/book"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Routes(router *gin.Engine)
}

type RouterContainer struct {
	bookHandler Router
	// Remaining handlers here...
}

func (r *RouterContainer) Routers() []Router {
	return []Router{r.bookHandler}
}

func NewRouterContainer(bookHandler *book.Handler) *RouterContainer {
	return &RouterContainer{bookHandler: bookHandler}
}
