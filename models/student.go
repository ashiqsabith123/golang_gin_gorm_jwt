package models

import "gorm.io/gorm"

type Students struct {
	gorm.Model
	Fname    string
	Lname    string
	Email    string
	Phone    string
	Place    string
	Dob      string
	Username string
	Password string
	Dep_id   string
}
