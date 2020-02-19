package models

//Survey the struct model of database
type Survey struct {
	ID       uint       `json:"id"`
	Title    string     `json:"title"`
	Question []Question `gorm:"auto_preload"`
}

//Question :nodoc:
type Question struct {
	ID       uint   `json:"-"`
	SurveyID uint   `json:"-"`
	Question string `json:"question"`
}
