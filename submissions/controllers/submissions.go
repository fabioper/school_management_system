package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fabioper/school_management_system/submissions/controllers/requests"
	"github.com/fabioper/school_management_system/submissions/models"
	"github.com/fabioper/school_management_system/submissions/services/contracts"
)

type SubmissionsController struct {
	database *gorm.DB
	service  contracts.ActivitiesService
}

func NewSubmissionsController(database *gorm.DB, service contracts.ActivitiesService) *SubmissionsController {
	return &SubmissionsController{database: database, service: service}
}

func (sc *SubmissionsController) GetAllSubmissions(c *gin.Context) {
	activityId, _ := strconv.Atoi(c.Query("activity_id"))
	var submissions []models.Submission
	sc.database.Where(&models.Submission{ActivityID: uint(activityId)}).Find(&submissions)
	c.JSON(http.StatusOK, submissions)
}

func (sc *SubmissionsController) SubmitActivity(c *gin.Context) {
	var request requests.SubmitActivityRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ficou doido é"})
		return
	}

	submission := sc.ParseSubmissionRequest(request)
	sc.database.Create(&submission)
	c.JSON(http.StatusCreated, submission)
}

func (sc *SubmissionsController) ParseSubmissionRequest(request requests.SubmitActivityRequest) models.Submission {
	submission := models.Submission{
		Content:     request.Content,
		ActivityID:  request.ActivityID,
		ClassroomID: request.ClassroomID,
		StudentID:   request.StudentID,
	}
	return submission
}

func (sc *SubmissionsController) FindSubmission(c *gin.Context) {
	id := c.Param("id")
	var submission models.Submission

	if err := sc.database.First(&submission, id).Error; err != nil {
		message := fmt.Sprintf("Submission with id #%v not found", id)
		c.JSON(http.StatusNotFound, gin.H{"error": message})
		return
	}

	c.JSON(http.StatusOK, submission)
}
