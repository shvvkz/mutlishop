package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}

func ConnectDatabase() {
	LoadEnv()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST"),
		getEnv("DB_USER"),
		getEnv("DB_PASSWORD"),
		getEnv("DB_NAME"),
		getEnv("DB_PORT"),
	)

	// 🔥 Logger personnalisé GORM
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Warn, // log seulement warnings & erreurs
			IgnoreRecordNotFoundError: true,        // ignore record not found
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	log.Println("Database connected")
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return value
}
