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
	if err := repo.DB.Find(&survey).Error; err != nil {
		return nil, err
	}
	return survey, nil
}

//GetAnswer get all question based on survey
func (repo *surveyRepo) GetAnswerRepo() ([]models.Answer, error) {
	var survey []models.Answer
	if err := repo.DB.Find(&survey).Error; err != nil {
		return nil, err
	}
	return survey, nil
}

//PostSubmit is function for buat bro
func (repo *surveyRepo) PostSubmitRepo(Answer models.Answer) (models.Answer, error) {
	if err := repo.DB.Create(&Answer).Error; err != nil {
		return models.Answer{}, err
	}
	return Answer, nil
}

//CreateSurvey :nodoc:
func (repo *surveyRepo) CreateSurveyRepo(id string) (models.Survey, error) {
	var data models.Survey
	if err := repo.DB.Where("id = ?", id).First(&data).Error; err != nil {
		fmt.Println(err)
		return models.Survey{}, err
	}
	return data, nil
}

//GetASurvey :nodoc:
func (repo *surveyRepo) GetASurveyRepo(id string) (models.Survey, error) {
	var data models.Survey
	if err := repo.DB.Where("id = ?", id).First(&data).Error; err != nil {
		fmt.Println(err)
		return models.Survey{}, err
	}
	return data, nil
}

//UpdateSurvey :nodoc:
func (repo *surveyRepo) UpdateSurveyRepo(data models.Survey, id string) (models.Survey, error) {
	repo.DB.Save(&data)
	return data, nil
}

//DeleteSurvey :nodoc:
func (repo *surveyRepo) DeleteSurveyRepo(data models.Survey, id string) (models.Survey, error) {
	if err := repo.DB.Where("id = ?", id).Delete(&data).Error; err != nil {
		return models.Survey{}, err
	}
	return data, nil
}
