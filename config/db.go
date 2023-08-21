package config

import (
	"login-register/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open("postgres", "host=91ca9cb7fd78 port=5432 user=pstgres dbname=loginregister sslmode=disable password=postgrespw")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Users{})
}
