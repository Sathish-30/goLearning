package database

import (
	"log"

	"github.com/sathish-30/gorilla-mux/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := "host=floppy.db.elephantsql.com user=xqhavbxn password=jmJoUTCnP14B6WRENOFZCQpI7TIURerx dbname=xqhavbxn port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database not connected")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(&model.Admin{})
}
