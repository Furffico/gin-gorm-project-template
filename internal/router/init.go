package router

import (
	"fmt"
	"log"
	"server/internal/config"

	"github.com/gin-gonic/gin"
)

func RunServer() (err error) {

	release := !config.Conf.General.Debug
	if release {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// load configurations
	r.SetTrustedProxies(config.Conf.Server.TrustedProxies)
	if release {
		r.Use(gin.Recovery())
	}

	// mount routers
	mountRouters(r)

	// mount static files
	mountStatic(r)

	// start server
	bind := fmt.Sprintf("%s:%d", config.Conf.Server.BindAddress, config.Conf.Server.BindPort)
	log.Printf("Starting server on http://%s/\n", bind)
	err = r.Run(bind)

	return
}
