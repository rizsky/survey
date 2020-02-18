package handlers

import (
	"fmt"
	"net/http"
	"soal_nomor_2/module/survey/models"
	"soal_nomor_2/module/survey/repositories"

	"github.com/gin-gonic/gin"
)

//Handler :nodoc:
type Handler struct {
	survey repositories.Repository
}

//NewSurveyHandler :nodoc:
func NewSurveyHandler(re repositories.Repository) *Handler {
	return &Handler{
		survey: re,
	}
}

// GetAllSurvey :nodoc:
func (repo *Handler) GetAllSurvey(c *gin.Context) {
	todo, err := repo.survey.GetAllSurveyRepo()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// GetAnswer :nodoc:
func (repo *Handler) GetAnswer(c *gin.Context) {
	_ = c.Params.ByName("id")
	todo, err := repo.survey.GetAnswerRepo()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// PostSubmit :nodoc:
func (repo *Handler) PostSubmit(c *gin.Context) {
	_ = c.Params.ByName("id")
	todo, err := repo.survey.PostSubmitRepo(models.Answer{})
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// GetASurvey :nodoc:
func (repo *Handler) GetASurvey(c *gin.Context) {
	id := c.Params.ByName("id")
	todo, err := repo.survey.GetASurveyRepo(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// CreateSurvey :nodoc:
func (repo *Handler) CreateSurvey(c *gin.Context) {
	var todo models.Survey
	requestPayload := make(map[string]interface{})

	err := c.BindJSON(&requestPayload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	fmt.Println(requestPayload)
	todo, err = repo.survey.CreateSurveyRepo("test")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

// EditSurvey :nodoc:
func (repo *Handler) EditSurvey(c *gin.Context) {
	// var todo models.Survey
	// id := c.Params.ByName("id")
	// todo, err := repo.survey.EditSurveyRepo(id)
	// if err != nil {
	// 	c.AbortWithError(http.StatusNotFound, err)
	// 	return
	// }
	// c.BindJSON(&todo)
	// todo, err = repo.survey.UpdateSurveyRepo(todo, id)
	// if err != nil {
	// 	c.AbortWithStatus(http.StatusNotFound)
	// } else {
	// 	c.JSON(http.StatusOK, todo)
	// }
}

// DeleteSurvey :nodoc:
func (repo *Handler) DeleteSurvey(c *gin.Context) {
	var todo models.Survey
	id := c.Params.ByName("id")
	todo, err := repo.survey.DeleteSurveyRepo(todo, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
