package main

import (
	"backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:52242"}

	router.Use(cors.New(config))
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/products", handlers.GetProducts)
	router.Run("localhost:8080")
}
