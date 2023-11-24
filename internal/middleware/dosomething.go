package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func DoSomethingMiddleware(c *gin.Context) {
	log.Println("Do something.")
	c.Next()
}
