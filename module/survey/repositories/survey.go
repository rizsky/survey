package repositories

import "soal_nomor_2/module/survey/models"

//Repository return contract
type Repository interface {
	GetAllSurveyRepo() ([]models.Survey, error)
	GetAnswerRepo(id uint) ([]models.Answer, error)
	PostSubmitRepo(surveyID uint, user string, answers []models.Answer) (models.Result, error)
	CreateSurveyRepo(name string, questions []string) (models.Survey, error)
	GetASurveyRepo(id string) (models.Survey, error)
	DeleteSurveyRepo(data models.Survey, id string) (models.Survey, error)
}
