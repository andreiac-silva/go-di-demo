package server

import (
	"net/http"
	"time"

	"github.com/andreiac-silva/go-di-demo/api"
	"github.com/andreiac-silva/go-di-demo/cmd/setup/env"

	"github.com/gin-gonic/gin"
)

func NewHTTPServerForWire(routerContainer *api.RouterContainer) *http.Server {
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
