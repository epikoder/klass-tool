package models

type User struct {
	Id       uint
	Username string `gorm:"unique"`
	Password string
	Fname    string
	Lname    string
	Result   uint
}
