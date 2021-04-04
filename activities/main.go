package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/fabioper/school_management_system/activities/docs"

	"github.com/fabioper/school_management_system/activities/controllers"
	"github.com/fabioper/school_management_system/activities/models"
)

// @title Activities Microservice
// @version 1.0
// @description Activities microservice for School Management System.
// @contact.name Fábio Pereira da Silva
// @contact.email fabio.dsilva@al.infnet.edu.br
// @license.name MIT
// @BasePath /api/
func main() {
	r := gin.Default()

	models.ConnectDatabase()

	ac := controllers.NewActivitiesController(models.DB)
	api := r.Group("/api")
	{
		api.GET("/activities", ac.GetAllActivities)
		api.POST("/activities", ac.PublishActivity)
		api.GET("/activities/:id", ac.FindActivity)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8000")
}
