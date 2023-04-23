package models

import "time"

type Subject struct {
	Id          uint   `sql:"primary_key;" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Questions   []Question
	Topics      []Topic
	Objectives  []Objective `json:"objectives"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
