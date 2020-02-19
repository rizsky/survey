package handlers

import (
	"net/http"
	"soal_nomor_2/module/survey/models"
	"soal_nomor_2/module/survey/repositories"

	"strconv"

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
	surveys, err := repo.survey.GetAllSurveyRepo()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, surveys)
	}
}

// GetAnswer :nodoc:
func (repo *Handler) GetAnswer(c *gin.Context) {
	idURL := c.Params.ByName("id")
	id, _ := strconv.Atoi(idURL)
	survey, err := repo.survey.GetAnswerRepo(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	var data []SubmissionAnswerResponse
	for _, answer := range survey {
		data = append(data, SubmissionAnswerResponse{
			UserName: answer.Result.UserName,
			Question: answer.Question.Question,
			Answer:   answer.Answer,
		})
	}

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

// PostSubmit :nodoc:
func (repo *Handler) PostSubmit(c *gin.Context) {
	var requestData SubmitSubmissionRequest
	id := c.Params.ByName("id")
	survey, err := repo.survey.GetASurveyRepo(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	err = c.BindJSON(&requestData)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	var answers []models.Answer
	for _, answerData := range requestData.Answers {
		answers = append(answers, models.Answer{
			QuestionID: answerData.QuestionID,
			Answer:     answerData.Answer,
		})
	}

	result, err := repo.survey.PostSubmitRepo(survey.ID, requestData.User, answers)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

// GetASurvey :nodoc:
func (repo *Handler) GetASurvey(c *gin.Context) {
	id := c.Params.ByName("id")
	survey, err := repo.survey.GetASurveyRepo(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, survey)
	}
}

// CreateSurvey :nodoc:
func (repo *Handler) CreateSurvey(c *gin.Context) {
	var (
		request CreateSurveyRequest
		survey  models.Survey
	)
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	survey, err = repo.survey.CreateSurveyRepo(request.Name, request.Questions)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, survey)
	}

}

// DeleteSurvey :nodoc:
func (repo *Handler) DeleteSurvey(c *gin.Context) {
	var survey models.Survey
	id := c.Params.ByName("id")
	survey, err := repo.survey.DeleteSurveyRepo(survey, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
