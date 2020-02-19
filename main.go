package main

import (
	"fmt"

	"soal_nomor_2/db"
	handler "soal_nomor_2/module/survey/handlers"
	"soal_nomor_2/module/survey/models"
	repo "soal_nomor_2/module/survey/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	db.DB, err = gorm.Open("mysql", db.DbURL(db.BuildDbConfig()))
	if err != nil {
		fmt.Println("status: ", err)
	}
	db.DB.AutoMigrate(&models.Answer{}, &models.Survey{}, &models.Result{}, &models.Question{})

	sr := repo.NewSurveyRepository(db.DB)
	controller := handler.NewSurveyHandler(sr)

	defer db.DB.Close()

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("survey", controller.GetAllSurvey)
		v1.GET("survey/:id", controller.GetASurvey)
		v1.GET("result/:id", controller.GetAnswer)
		v1.DELETE("survey/:id", controller.DeleteSurvey)
		v1.POST("survey", controller.CreateSurvey)
		v1.POST("submit/:id", controller.PostSubmit)
	}

	//running server
	r.Run()
}
