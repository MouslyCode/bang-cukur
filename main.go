package main

import (
	"log"

	"github.com/MouslyCode/bang-cukur/database"
	"github.com/MouslyCode/bang-cukur/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	database.Connect()
	r := gin.Default()
	routes.AuthRoutes(r)
	routes.ItemRoutes(r)
	routes.TransactionRoutes(r)

	r.Run(":8080")
}
