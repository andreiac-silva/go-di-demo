package server

import (
	"net/http"
	"time"

	"github.com/andreiac-silva/go-di-demo/api"
	"github.com/andreiac-silva/go-di-demo/cmd/setup/env"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(routers []api.Router) *http.Server {
	r := gin.New()

	for _, handler := range routers {
		handler.Routes(r)
	}

	return &http.Server{
		Addr:         env.ServerAddress,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
}
