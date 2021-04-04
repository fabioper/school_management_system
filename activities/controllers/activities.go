package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	. "github.com/fabioper/school_management_system/activities/controllers/requests"
	"github.com/fabioper/school_management_system/activities/models"
)

func GetAllActivities(c *gin.Context) {
	var books []models.Activity
	models.DB.Find(&books)

	c.JSON(http.StatusOK, books)
}

func AddActivity(c *gin.Context) {
	var request CreateActivityRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	activity := models.Activity{Content: request.Content}
	models.DB.Create(&activity)
	c.JSON(http.StatusCreated, activity)
}

func FindActivity(c *gin.Context) {
	id := c.Param("id")
	var activity models.Activity

	if err := models.DB.First(&activity, id).Error; err != nil {
		message := fmt.Sprintf("Activity with id #%v not found", id)
		c.JSON(http.StatusNotFound, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, activity)
}
