package router

import (
	"log"
	"server/internal/config"
	m "server/internal/middleware"
	s "server/internal/service"

	"github.com/gin-gonic/gin"
)

func mountRouters(r *gin.Engine) {
	basepath := config.Conf.Server.APIBasePath
	root := r.Group(basepath)

	root.GET("/ping", m.DoSomethingMiddleware, s.Ping)

	log.Println("Routers mounted")
}
