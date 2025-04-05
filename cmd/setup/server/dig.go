package server

import (
	"net/http"
	"time"

	"github.com/andreiac-silva/go-di-demo/api"
	"github.com/andreiac-silva/go-di-demo/cmd/setup/env"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Routers struct {
	dig.In
	Routers []api.Router `group:"routers"`
}

func NewHTTPServerForDig(in Routers) *http.Server {
	r := gin.New()

	for _, handler := range in.Routers {
		handler.Routes(r)
	}

	return &http.Server{
		Addr:         env.ServerAddress,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
}
