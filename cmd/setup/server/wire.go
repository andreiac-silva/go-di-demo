package server

import (
	"net/http"
	"time"

	"github.com/andreiac-silva/go-di-demo/api"
	"github.com/andreiac-silva/go-di-demo/book"
	"github.com/andreiac-silva/go-di-demo/cmd/setup/env"

	"github.com/gin-gonic/gin"
)

// https://github.com/google/wire/issues/207

type RouterContainer struct {
	bookHandler api.Router
	// Remaining handlers here or a slice of routers...
}

func (r *RouterContainer) Routers() []api.Router {
	return []api.Router{r.bookHandler}
}

func NewRouterContainer(bookHandler *book.Handler) *RouterContainer {
	return &RouterContainer{bookHandler: bookHandler}
}

func NewHTTPServerForWire(routerContainer *RouterContainer) *http.Server {
	r := gin.Default()

	for _, handler := range routerContainer.Routers() {
		handler.Routes(r)
	}

	return &http.Server{
		Addr:         env.ServerAddress,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
}
