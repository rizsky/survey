package models

//Survey the struct model of database
type Survey struct {
	ID       uint       `json:"id"`
	Title    string     `json:"title"`
	Question []Question `gorm:"-"`
}

//Question :nodoc:
type Question struct {
	ID       uint   `json:"id"`
	IDSurvey uint   `json:id_survey`
	Question string `json:"question"`
}
