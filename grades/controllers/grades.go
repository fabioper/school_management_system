package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GradesController struct {
	database *gorm.DB
}

func NewGradesController(database *gorm.DB) *GradesController {
	return &GradesController{database: database}
}

func (gc *GradesController) GetGrades(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (gc *GradesController) PublishGrade(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}

