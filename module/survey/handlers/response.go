package handlers

//SubmissionAnswerResponse :nodoc:
type SubmissionAnswerResponse struct {
	UserName string   `json:"user_name"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
