package models

type (
	Objective struct {
		Id        uint                `sql:"primary_key;" json:"id"`
		SubjectId int                 `json:"subject_id"`
		Title     string              `json:"title"`
		Details   string              `json:"details"`
		Questions []ObjectiveQuestion `json:"questions"`
	}

	ObjectiveQuestion struct {
		Id          uint `sql:"primary_key;" json:"id"`
		ObjectiveId int  `json:"objective_id"`
		QuestionId  int  `json:"question_id"`
	}

	TopicObjective struct {
		Id          uint   `sql:"primary_key;" json:"id"`
		TopicId     int    `json:"topic_id"`
		ObjectiveId int    `json:"objective_id"`
		Title       string `json:"title"`
		Details     string `json:"details"`
	}
)
