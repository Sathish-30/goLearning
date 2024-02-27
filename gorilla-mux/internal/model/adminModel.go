package model

import "gorm.io/gorm"

type Admin struct {
	AdminName string
	AdminAge  int
	gorm.Model
}
