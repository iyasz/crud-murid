package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnetDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_murid"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Class{}, &Murid{})
	DB = db
}
