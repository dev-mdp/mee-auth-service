package main

import (
	"authservice/src/api"
	"authservice/src/config"
	"authservice/src/models"
)

func main() {
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.MasterRoles{})
	config.DB.AutoMigrate(&models.MasterStatus{})

	api.APIRoutes()
}
