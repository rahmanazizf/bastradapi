package main

import (
	"basic-trade-api/database"
	"basic-trade-api/routers"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// var PORT = os.Getenv("PORT")

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env file:", err)
		return // Keluar dari program jika gagal memuat file .env
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		fmt.Println("PORT is not set in .env file")
		return // Keluar dari program jika variabel PORT tidak diatur
	}

	database.DBConnection()
	// fmt.Println(fmt.Sprintf("Listening on :%s", PORT))
	fmt.Printf("Listening on :%s\n", PORT)
	routers.StartServer().Run(fmt.Sprintf(":%s", PORT))
}
