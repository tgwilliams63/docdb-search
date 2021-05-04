package api_endpoints

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func CreateRouter() *gin.Engine {
	log.Printf("Gin cold start")
	r := gin.Default()
	r.Use(cors.Default())

	//r.LoadHTMLGlob("web/templates/*")
	r.GET("/fields", GetFields)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}