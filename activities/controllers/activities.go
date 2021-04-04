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

func (ac *ActivitiesController) GetAllActivities(c *gin.Context) {
	var books []models.Activity
	ac.database.Find(&books)

	c.JSON(http.StatusOK, books)
}

func (ac *ActivitiesController) AddActivity(c *gin.Context) {
	var request CreateActivityRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	activity := models.Activity{Content: request.Content}
	ac.database.Create(&activity)
	c.JSON(http.StatusCreated, activity)
}

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
