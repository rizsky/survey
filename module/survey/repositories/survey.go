package repositories

import "soal_nomor_2/module/survey/models"

//Repository return contract
type Repository interface {
	GetAllSurveyRepo() ([]models.Survey, error)
	GetAnswerRepo() ([]models.Answer, error)
	PostSubmitRepo(Answer models.Answer) (models.Answer, error)
	CreateSurveyRepo(id string) (models.Survey, error)
	GetASurveyRepo(id string) (models.Survey, error)
	UpdateSurveyRepo(data models.Survey, id string) (models.Survey, error)
	DeleteSurveyRepo(data models.Survey, id string) (models.Survey, error)
}
