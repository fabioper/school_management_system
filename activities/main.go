package main

import (
	"log"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/fabioper/school_management_system/activities/docs"

	"github.com/fabioper/school_management_system/activities/controllers"
	"github.com/fabioper/school_management_system/activities/models"
	. "github.com/fabioper/school_management_system/activities/services"
)

// @title Activities - School Management System
// @version 1.0
// @description Service for managing activities.
// @termsOfService http://swagger.io/terms/

// @contact.name Fábio Pereira da Silva
// @contact.email fabio.dsilva@al.infnet.edu.br

// @BasePath /api
func main() {
	r := gin.Default()

	models.ConnectDatabase()

	ac := controllers.NewActivitiesController(models.DB, NewSubmissionsService())
	api := r.Group("/api")
	{
		api.GET("/activities", ac.GetAllActivities)
		api.POST("/activities", ac.PublishActivity)
		api.GET("/activities/:id", ac.FindActivity)
		api.GET("/activities/:id/submissions", ac.GetActivitySubmission)
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	log.Fatal(r.Run(":8000"))
}