package handlers

//CreateSurveyRequest :nodoc:
type CreateSurveyRequest struct {
	Name      string   `json:"name"`
	Questions []string `json:"questions"`
}

// SubmitSubmissionRequest :nodoc:
type SubmitSubmissionRequest struct {
	User    string                    `json:"user"`
	Answers []SubmissionAnswerRequest `json:"answers"`
}

// SubmissionAnswerRequest :nodoc:
type SubmissionAnswerRequest struct {
	QuestionID uint   `json:"question_id"`
	Answer     string `json:"answer"`
}
