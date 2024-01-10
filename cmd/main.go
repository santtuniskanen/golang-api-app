package main

import (
	"github.com/santtuniskanen/golang-restapi/database"
	"github.com/santtuniskanen/golang-restapi/router"
)

func init() {
	database.LoadEnvironment()
	database.ConnectToDB()
}

func main() {
	router.Run()
}
