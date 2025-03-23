package setup

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(routers ...Router) *http.Server {
	r := gin.Default()

	for _, handler := range routers {
		handler.Routes(r)
	}

	a := os.Getenv("SERVER_ADDRESS")

	return &http.Server{
		Addr:         a,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
}

type Router interface {
	Routes(router *gin.Engine)
}
