package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	connectDatabase()
}

func connectDatabase() {
	host := "localhost"
	port := "5432"
	database := "Whoa"
	username := "postgres"
	password := "postgres"
	timezone := "Asia/Shanghai"
	dsn := "host=" + host + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable TimeZone=" + timezone
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database, err: " + err.Error())
	}
	DB.AutoMigrate(&User{})
}
