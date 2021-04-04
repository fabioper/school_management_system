package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/fabioper/school_management_system/activities/controllers"
	"github.com/fabioper/school_management_system/activities/models"
)

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

	log.Fatal(r.Run(":8000"))
}