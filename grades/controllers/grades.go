package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	. "github.com/fabioper/school_management_system/grades/controllers/requests"
	"github.com/fabioper/school_management_system/grades/models"
	"github.com/fabioper/school_management_system/grades/services"
)

type GradesController struct {
	database *gorm.DB
}

func NewGradesController(database *gorm.DB) *GradesController {
	return &GradesController{database: database}
}

func (gc *GradesController) GetGrades(c *gin.Context) {
	var grades []models.Grade
	gc.database.Find(&grades)
	c.JSON(http.StatusOK, grades)
}

func (gc *GradesController) PublishGrade(c *gin.Context) {
	var request PublishGradeRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	newGrade := gc.ToModel(request)
	gc.database.Create(&newGrade)
	services.PublishMessage("grade-published", newGrade)

	c.JSON(http.StatusCreated, newGrade)
}

func (gc *GradesController) ToModel(request PublishGradeRequest) models.Grade {
	newGrade := models.Grade{
		Model:        gorm.Model{},
		SubmissionID: request.SubmissionID,
		TeacherID:    request.TeacherID,
		Grade:        request.Grade,
	}
	return newGrade
}
