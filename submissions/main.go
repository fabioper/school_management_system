package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/fabioper/school_management_system/submissions/controllers"
	"github.com/fabioper/school_management_system/submissions/models"
	. "github.com/fabioper/school_management_system/submissions/services"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	sc := controllers.NewSubmissionsController(models.DB, NewExternalActivitiesService())
	api := r.Group("/api")
	{
		api.GET("/submissions", sc.GetAllSubmissions)
		api.POST("/submissions", sc.SubmitActivity)
		api.GET("/submissions/:id", sc.FindSubmission)
	}

	log.Fatal(r.Run(":5000"))
}
