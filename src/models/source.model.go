package models

import "time"

type Source struct {
	Id        uint      `sql:"primary_key;" json:"-"`
	Name      string    `json:"-"`
	Level     string    `json:"-"`
	Year      uint      `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
