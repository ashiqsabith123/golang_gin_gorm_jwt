package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string
	Username string
	Password string
}
