package repositories

import (
	"fmt"
	"soal_nomor_2/module/survey/models"

	_ "github.com/go-sql-driver/mysql" //mysql driver
	"github.com/jinzhu/gorm"
)

type surveyRepo struct {
	DB *gorm.DB
}

//NewSurveyRepository return new instance
func NewSurveyRepository(db *gorm.DB) Repository {
	return &surveyRepo{
		DB: db,
	}
}

//GetAllSurvey dapetin semua survey
func (repo *surveyRepo) GetAllSurveyRepo() ([]models.Survey, error) {
	var survey []models.Survey
	if err := repo.DB.Preload("Question").Find(&survey).Error; err != nil {
		return nil, err
	}

	return survey, nil
}

//GetAnswer get all question based on survey
func (repo *surveyRepo) GetAnswerRepo(id uint) ([]models.Answer, error) {
	var answer []models.Result
	repo.DB.Preload("Answers").Preload("Answers.Result").Preload("Answers.Question").Find(&answer)

	var submitAnswer []models.Answer
	for _, submit := range answer {
		submitAnswer = append(submitAnswer, submit.Answers...)
	}

	return submitAnswer, nil
}

//PostSubmit is function for buat bro
func (repo *surveyRepo) PostSubmitRepo(surveyID uint, user string, answers []models.Answer) (models.Result, error) {
	data := models.Result{
		SurveyID: surveyID,
		UserName: user,
		Answers:  answers,
	}

	if err := repo.DB.Debug().Save(&data).Error; err != nil {
		return models.Result{}, err
	}
	return data, nil
}

//CreateSurvey :nodoc:
func (repo *surveyRepo) CreateSurveyRepo(title string, questions []string) (models.Survey, error) {
	var setquestions []models.Question
	for _, question := range questions {
		setquestions = append(setquestions, models.Question{
			Question: question,
		})
	}

	survey := models.Survey{
		Title:    title,
		Question: setquestions,
	}

	if err := repo.DB.Save(&survey).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

//GetASurvey :nodoc:
func (repo *surveyRepo) GetASurveyRepo(id string) (models.Survey, error) {
	var data models.Survey
	if err := repo.DB.Preload("Question").Where("id = ?", id).First(&data).Error; err != nil {
		fmt.Println(err)
		return models.Survey{}, err
	}
	return data, nil
}

//DeleteSurvey :nodoc:
func (repo *surveyRepo) DeleteSurveyRepo(data models.Survey, id string) (models.Survey, error) {
	if err := repo.DB.Where("id = ?", id).Delete(&data).Error; err != nil {
		return models.Survey{}, err
	}
	return data, nil
}
