package handlers

import (
	"backend/data"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	c.JSON(200, data.Products)
}
