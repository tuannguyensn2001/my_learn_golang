package gincourse

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Index(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": "o trong index",
		})
	}
}
