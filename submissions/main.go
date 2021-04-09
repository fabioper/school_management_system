package main

import (
	"log"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/fabioper/school_management_system/submissions/docs"

	"github.com/fabioper/school_management_system/submissions/controllers"
	"github.com/fabioper/school_management_system/submissions/models"
	. "github.com/fabioper/school_management_system/submissions/services"
	"github.com/fabioper/school_management_system/submissions/services/consumers"
)

// @title Submissions - School Management System
// @version 1.0
// @description Service for managing submissions.
// @contact.name Fábio Pereira da Silva
// @contact.email fabio.dsilva@al.infnet.edu.br
// @BasePath /api
func main() {
	r := gin.Default()

	go consumers.StartConsumers()
	models.ConnectDatabase()

	sc := controllers.NewSubmissionsController(models.DB, NewExternalActivitiesService())
	api := r.Group("/api")
	{
		api.GET("/submissions", sc.GetAllSubmissions)
		api.POST("/submissions", sc.SubmitActivity)
		api.GET("/submissions/:id", sc.FindSubmission)
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	log.Fatal(r.Run(":5000"))
}
