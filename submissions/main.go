package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/fabioper/school_management_system/submissions/controllers"
	"github.com/fabioper/school_management_system/submissions/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	sc := controllers.NewSubmissionsController(models.DB)
	api := r.Group("/api")
	{
		api.GET("/submissions", sc.GetAllSubmissions)
		api.POST("/submissions", sc.AddSubmission)
		api.GET("/submissions/:id", sc.FindSubmission)
	}

	log.Fatal(r.Run(":5000"))
}
