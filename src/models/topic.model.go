package models

type Topic struct {
	Id         uint             `sql:"primary_key;" json:"id"`
	SubjectId  int              `json:"subject_id"`
	Title      string           `json:"title"`
	Details    string           `json:"details"`
	Objectives []TopicObjective `json:"objectives"`
}
