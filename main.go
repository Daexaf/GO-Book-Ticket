package main

import (
	"example.com/BookEvent/routes"

	"example.com/BookEvent/db"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
