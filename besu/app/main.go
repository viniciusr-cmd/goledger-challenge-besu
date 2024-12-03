package main

import (
	"goledger-challenge/vinicius/besu/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	r := gin.Default()
	routes.InitRoutes(r)
	r.Run(":8080")
}
