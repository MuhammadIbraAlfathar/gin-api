package main

import (
	"fmt"
	"github.com/MuhammadIbraAlfathar/gin-api/config"
	"github.com/MuhammadIbraAlfathar/gin-api/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	config.LoadConfig()
	config.LoadDB()

	r := gin.Default()
	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.AuthRouter(api)

	err := r.Run(fmt.Sprintf(":%v", config.ENV.PORT))
	if err != nil {
		log.Fatal("Error connection to port")
	} // listen and serve on 0.0.0.0:8080
}
