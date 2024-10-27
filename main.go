package main

import (
	"github.com/titi0001/Go-Gin-RestAPI/database"
	"github.com/titi0001/Go-Gin-RestAPI/routes"
)

func main() {
	database.ConectaDb()
	routes.HandleRequest()
}
