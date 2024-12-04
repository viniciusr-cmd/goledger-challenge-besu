package main

import (
	"goledger-challenge/vinicius/besu/routes"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var db *gorm.DB

type Contract struct {
	ID        uint `gorm:"primaryKey"`
	Value     int64
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("POSTGRES_URL")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	err = db.AutoMigrate(&Contract{})
	if err != nil {
		log.Fatalf("Error migrating the database: %v", err)
	}

	initialContract := Contract{
		Address: os.Getenv("CONTRACT_ADDRESS"),
		Value:   0,
	}

	if err := db.FirstOrCreate(&initialContract, Contract{Address: initialContract.Address}).Error; err != nil {
		log.Fatalf("Error creating initial contract record: %v", err)
	}
}

func main() {
	r := gin.Default()
	routes.InitRoutes(r, db)
	r.Run(":8080")
}
