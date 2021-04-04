package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fabioper/school_management_system/submissions/controllers/requests"
	"github.com/fabioper/school_management_system/submissions/models"
)

type SubmissionsController struct {
	database *gorm.DB
}

func (sc *SubmissionsController) GetAllSubmissions(c *gin.Context) {
	var books []models.Submission
	sc.database.Find(&books)
	c.JSON(http.StatusOK, books)
}

func (sc *SubmissionsController) AddSubmission(c *gin.Context) {
	var request requests.AddSubmissionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": "ficou doido é" })
		return
	}

	submission := sc.ParseSubmissionRequest(request)
	sc.database.Create(&submission)
	c.JSON(http.StatusCreated, submission)
}

func (sc *SubmissionsController) ParseSubmissionRequest(request requests.AddSubmissionRequest) models.Submission {
	submission := models.Submission{
		Content:     request.Content,
		ActivityId:  request.ActivityId,
		ClassroomId: request.ClassroomId,
		StudentId:   request.StudentId,
	}
	return submission
}

func (sc *SubmissionsController) FindSubmission(c *gin.Context) {
	// TODO
}

func NewSubmissionsController(database *gorm.DB) *SubmissionsController {
	return &SubmissionsController{database: database}
}