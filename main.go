package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	coursesroute "my_learn_golang/modules/courses"
)

func main() {
	dsn := "root:password@tcp(172.17.0.1:3306)/golang?charset=utf8&parseTime=True"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	r := gin.Default()

	coursesroute.HandleRoute(r,db)

	err := r.Run(":6000")
	if err != nil {
		return
	}

}
