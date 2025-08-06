package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load("src/config/.env")
	if err != nil {
		log.Fatal("Error loading .env file : ", err)
	}

	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	DBPort := os.Getenv("DB_PORT")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBSSLMode := os.Getenv("DB_SSLMODE")
	DBTimeZone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		DBHost, DBUser, DBPassword, DBName, DBPort, DBSSLMode, DBTimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("cannot connect to database : " + err.Error())
		log.Fatal("Error connecting to database : ", err)
	} else {
		fmt.Println("connected to database")
	}

	DB = db
}
