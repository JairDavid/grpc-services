package domain

type Question struct {
	Id       string `json:"id"`
	Answer   string `json:"answer"`
	Question string `json:"question"`
	ExamId   string `json:"exam_id"`
}
