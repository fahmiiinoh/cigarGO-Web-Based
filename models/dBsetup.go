package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(mysql.Open("fahmidevtest:tmp_pwd@tcp(127.0.0.1:3306)/cigargo?charset=utf8&parseTime=True"), &gorm.Config{})

	if err != nil {
		panic(("Failed to connect to database!"))
	}
	DB = database
}

func DBMigrate() {

	DB.AutoMigrate(&Smoker{})

}
