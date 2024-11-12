// package config

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func LoadEnv() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

// func SetupDB() *gorm.DB {
// 	LoadEnv()

// 	host := os.Getenv("DB_HOST")
// 	port := os.Getenv("DB_PORT")
// 	user := os.Getenv("DB_USER")
// 	password := os.Getenv("DB_PASSWORD")
// 	dbname := os.Getenv("DB_NAME")
// 	schema := os.Getenv("DB_SCHEMA")

// 	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=disable", host, port, user, password, dbname, schema)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("failed to connect to database:", err)
// 	}
// 	return db
// }
