package models

//Answer the struct model of database
type Answer struct {
	ID         uint     `json:"id"`
	QuestionID uint     `json:"-"`
	ResultID   uint     `json:"-"`
	Answer     string   `json:"answer"`
	Question   Question `json:"question" gorm:"foreignkey:QuestionID"`
	Result     Result   `json:"result" gorm:"foreignkey:ResultID"`
}

//Result :nodoc:
type Result struct {
	ID       uint     `json:"id"`
	SurveyID uint     `json:"id_survey"`
	UserName string   `json:"user_name"`
	Survey   Survey   `json:"survey"`
	Answers  []Answer `json:"answers"`
}
