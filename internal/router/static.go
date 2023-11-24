package router

import (
	"log"
	"server/internal/config"

	"github.com/gin-gonic/gin"
)

func mountStatic(r *gin.Engine) {
	conf := config.Conf.Server.Static
	if conf.Enabled {
		r.Static(conf.BasePath, conf.Location)
		r.StaticFile("/", conf.IndexPage)
		log.Println("Static files mounted")
	}
}
