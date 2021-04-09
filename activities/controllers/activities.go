package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	. "github.com/fabioper/school_management_system/activities/controllers/requests"
	"github.com/fabioper/school_management_system/activities/models"
	. "github.com/fabioper/school_management_system/activities/services/contracts"
)

type ActivitiesController struct{
	database           *gorm.DB
	submissionsService SubmissionsService
}

func NewActivitiesController(db *gorm.DB, submissionsService SubmissionsService) *ActivitiesController {
	return &ActivitiesController{
		database: db,
		submissionsService: submissionsService,
	}
}

// GetAllActivities godoc
// @Summary Busca atividades
// @Produce json
// @Success 200 {array} models.Activity
// @Router /activities/ [get]
func (ac *ActivitiesController) GetAllActivities(c *gin.Context) {
	var activities []models.Activity
	ac.database.Find(&activities)

	c.JSON(http.StatusOK, activities)
}

// PublishActivity godoc
// @Summary Publica uma nova atividade
// @Accept  json
// @Produce  json
// @Param activity body PublishActivityRequest true "Activity"
// @Success 201 {object} models.Activity
// @Router /activities/ [post]
func (ac *ActivitiesController) PublishActivity(c *gin.Context) {
	var request PublishActivityRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestDeadline, err := time.Parse("02/01/2006 15:04", request.Deadline)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	activity := models.Activity{
		Content:   request.Content,
		TeacherID: request.TeacherID,
		Deadline:  requestDeadline,
	}

	ac.database.Create(&activity)
	c.JSON(http.StatusCreated, activity)
}

// FindActivity godoc
// @Summary Busca uma atividade pela ID
// @Produce json
// @Success 200 {object} models.Activity
// @Router /activities/{id} [get]
func (ac *ActivitiesController) FindActivity(c *gin.Context) {
	id := c.Param("id")
	var activity models.Activity

	if err := ac.database.First(&activity, id).Error; err != nil {
		message := fmt.Sprintf("Activity with id #%v not found", id)
		c.JSON(http.StatusNotFound, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, activity)
}

// GetActivitySubmission godoc
// @Summary Busca as submissões de uma atividade
// @Produce json
// @Success 200 {array} Submission
// @Router /activities/{id}/submissions [get]
func (ac *ActivitiesController) GetActivitySubmission(c *gin.Context) {
	id := c.Param("id")
	submissions, err := ac.submissionsService.GetSubmissions(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, submissions)
}