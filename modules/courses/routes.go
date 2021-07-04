package coursesroute

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	gincourse "my_learn_golang/modules/courses/transport/gin"
)

func HandleRoute(r *gin.Engine, db *gorm.DB) {
	r.GET("/courses", gincourse.Index(db))
	r.POST("/courses", gincourse.Create(db))
}
