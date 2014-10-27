package main

import (
	"love_golang/dbWrapper"
	"love_golang/models"
)

func performMigrations() {
	db := dbWrapper.DB()
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.User{})
}
