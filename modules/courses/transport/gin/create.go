package gincourse

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	coursemodel "my_learn_golang/modules/courses/model"
	courseservice "my_learn_golang/modules/courses/service"
	coursestorage "my_learn_golang/modules/courses/storage"
	"net/http"
)

func Create(db *gorm.DB) func(context *gin.Context) {
	return func(context *gin.Context) {

		var newCourse coursemodel.CourseCreate

		if err := context.ShouldBind(&newCourse); err !=  nil{
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "xin loi",
			})
		}

		store := coursestorage.NewSQLStore(db)

		service := courseservice.NewCreateCourseService(store)

		if err := service.Handle(context,&newCourse); err != nil {
			context.JSON(http.StatusBadRequest,gin.H{
				"message": "xin loi",
			})
		}

		context.JSON(http.StatusOK, gin.H{
			"data": newCourse,
		})
	}
}
