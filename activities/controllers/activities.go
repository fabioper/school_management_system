package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	. "github.com/fabioper/school_management_system/activities/controllers/requests"
	"github.com/fabioper/school_management_system/activities/models"
)

type ActivitiesController struct{
	database *gorm.DB
}

func NewActivitiesController(db *gorm.DB) *ActivitiesController {
	return &ActivitiesController{database: db}
}

// GetAllActivities godoc
// @Summary Retrieves all activities
// @Produce json
// @Success 200 {array} models.Activity
// @Router /activities [get]
func (ac *ActivitiesController) GetAllActivities(c *gin.Context) {
	var books []models.Activity
	ac.database.Find(&books)

	c.JSON(http.StatusOK, books)
}

// PublishActivity godoc
// @Summary Publish a new activity
// @Produce json
// @Accept json
// @Success 201
// @Param activity body PublishActivityRequest true "Publish activity"
// @Router /activities [post]
func (ac *ActivitiesController) PublishActivity(c *gin.Context) {
	var request PublishActivityRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	activity := models.Activity{Content: request.Content}
	ac.database.Create(&activity)
	c.JSON(http.StatusCreated, activity)
}

// FindActivity godoc
// @Summary Retrivies an activity by the given ID
// @Produce json
// @Accept json
// @Success 200 {object} models.Activity
// @Param id path int true "Activity ID"
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
