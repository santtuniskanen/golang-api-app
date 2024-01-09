package main

import (
	"github.com/santtuniskanen/golang-restapi/database"
	"github.com/santtuniskanen/golang-restapi/models"
)

func init() {
	database.LoadEnvironment()
	database.ConnectToDB()
}

func main() {
	database.DB.AutoMigrate(&models.Post{})
}
