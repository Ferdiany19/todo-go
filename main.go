package main

import (
	"todos-api/src/config"
	"todos-api/src/routes"

	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB = config.ConnectDB()
	defer config.DisconnectDB(db)

	routes.Routes()
}
