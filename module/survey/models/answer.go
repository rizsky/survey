package models

//Answer the struct model of database
type Answer struct {
	ID         uint   `json:"id"`
	IDQuestion uint   `json:"question_id"`
	IDUser     uint   `json:"user_id"`
	Answer     string `json:"answer"`
}

//Result :nodoc:
type Result struct {
	ID       uint       `json:"id"`
	IDSurvey uint       `json:"id_survey"`
	UserID   uint       `json:"user_id"`
	Question []Question `gorm:"-"`
	Answer   []Answer   `gorm:="-"`
}
