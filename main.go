package main

import (
	"fmt"
	"log"

	"github.com/MouslyCode/bang-cukur/common/helper"
	"github.com/MouslyCode/bang-cukur/database"
	"github.com/MouslyCode/bang-cukur/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	hash, _ := helper.HashPassword("jokiyakin22")
	fmt.Println(hash)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	database.Connect()
	r := gin.Default()
	routes.AuthRoutes(r)

	r.Run(":8080")
}
