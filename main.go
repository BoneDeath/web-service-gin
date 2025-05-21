package main

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/products", handlers.GetProducts)
	router.Run("localhost:8080")
}
