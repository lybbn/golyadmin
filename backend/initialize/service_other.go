//go:build !windows
// +build !windows

package initialize

import (
	"time"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

type Server interface {
	ListenAndServe() error
}

func InitServer(address string, router *gin.Engine) Server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
