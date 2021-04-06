package main

import (
	"github.com/gin-gonic/gin"

	"github.com/fabioper/school_management_system/grades/controllers"
	"github.com/fabioper/school_management_system/grades/models"
)

func main() {
	r := gin.Default()

	if err := models.ConnectDatabase(); err != nil {
		panic(err.Error())
	}

	gc := controllers.NewGradesController(models.DB)

	api := r.Group("/api")
	{
		api.GET("/grades", gc.GetGrades)
		api.POST("/grades", gc.PublishGrade)
	}
}