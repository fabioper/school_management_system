package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fabioper/school_management_system/submissions/controllers/requests"
	. "github.com/fabioper/school_management_system/submissions/controllers/responses"
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
	var books []models.Submission
	sc.database.Find(&books)
	c.JSON(http.StatusOK, books)
}

func (sc *SubmissionsController) AddSubmission(c *gin.Context) {
	var request requests.AddSubmissionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ficou doido é"})
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
	id := c.Param("id")
	var submission models.Submission

	if err := sc.database.First(&submission, id).Error; err != nil {
		message := fmt.Sprintf("Submission with id #%v not found", id)
		c.JSON(http.StatusNotFound, gin.H{"error": message})
		return
	}

	var details SubmissionDetailsResponse
	if err := sc.appendActivity(submission, &details); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, details)
}

func (sc *SubmissionsController) appendActivity(submission models.Submission, response *SubmissionDetailsResponse) error {
	activity, err := sc.service.FetchActivity(submission.ActivityId)
	if err != nil {
		return err
	}

	*response = SubmissionDetailsResponse{
		Content:     submission.Content,
		ClassroomId: submission.ID,
		StudentId:   submission.StudentId,
		Grade:       submission.Grade,
		Activity: ActivityContent{
			Id:      activity.Id,
			Content: activity.Content,
		},
	}
	return nil
}
