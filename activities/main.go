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

	api := r.Group("/api")
	{
		api.GET("/activities", controllers.GetAllActivities)
		api.POST("/activities", controllers.AddActivity)
		api.GET("/activities/:id", controllers.FindActivity)
	}

	log.Fatal(r.Run(":8000"))
}