package model

import "gorm.io/gorm"

type Admin struct {
	AdminName string `json:"adminName"`
	AdminAge  int    `json:"adminAge"`
	gorm.Model
}
