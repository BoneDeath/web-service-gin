package handlers

import (
	"backend/data"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.JSON(200, data.Albums)
}
