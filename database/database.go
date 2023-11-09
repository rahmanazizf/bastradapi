package database

import (
	"basic-trade-api/helpers"
	"basic-trade-api/models"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	var (
		host     = os.Getenv("HOST")
		port, _  = strconv.Atoi(os.Getenv("PORT"))
		user     = os.Getenv("USER")
		password = os.Getenv("PASSWORD")
		dbname   = os.Getenv("DBNAME")
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.CheckError(err)

	db.Exec("DROP TABLE IF EXISTS variants")
	db.Exec("DROP TABLE IF EXISTS products")
	db.Exec("DROP TABLE IF EXISTS admins")
	log.Println("Dropped existing tables")

	db.AutoMigrate(&models.Admin{}, &models.Product{}, &models.Variant{})

	return db

}
