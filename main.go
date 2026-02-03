package main

import (
	"log"

	"github.com/MouslyCode/bang-cukur/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	database.Connect()

}
