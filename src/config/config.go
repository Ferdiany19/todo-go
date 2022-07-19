package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dbHost := "localhost"
	dbUser := "root"
	dbPassword := "root"
	dbName := "go_todo"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:8889)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPassword, dbHost, dbName)

	db, errorDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect to db")
	}
	return db
}

func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}

	dbSQL.Close()
}
