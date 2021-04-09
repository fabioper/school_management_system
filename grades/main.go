package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/fabioper/school_management_system/grades/docs"

	"github.com/fabioper/school_management_system/grades/controllers"
	"github.com/fabioper/school_management_system/grades/models"
)

// @title Grades - School Management System
// @version 1.0
// @description Service for managing grades.
// @contact.name Fábio Pereira da Silva
// @contact.email fabio.dsilva@al.infnet.edu.br
// @BasePath /api
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
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	log.Fatal(r.Run(":8080"))
}