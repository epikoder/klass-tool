package models

import "time"

type Question struct {
	Id                 *uint     `sql:"primary_key;" json:"id"`
	SubjectId          uint      `json:"subject_id"`
	QuestionTypeId     uint      `json:"question_type_id"`
	Question           string    `json:"question"`
	QuestionDetails    string    `json:"question_details"`
	QuestionImage      string    `json:"question_image"`
	Option_1           string    `json:"option_1"`
	Option_2           string    `json:"option_2"`
	Option_3           string    `json:"option_3"`
	Option_4           string    `json:"option_4"`
	ShortAnswer        string    `json:"short_answer"`
	FullAnswer         string    `json:"full_answer"`
	AnswerImage        string    `json:"answer_image"`
	AnswerDetails      string    `json:"answer_details"`
	QuestionYear       uint      `json:"question_year"`
	QuestionYearNumber uint      `json:"question_year_number"`
	CreatedAt          time.Time `json:"-"`
	UpdatedAt          time.Time `json:"-"`
}
