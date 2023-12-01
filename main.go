package main

import (
	"basic-trade-api/database"
	"basic-trade-api/routers"
	"fmt"
)

func main() {
	database.DBConnection()
	fmt.Println("Listening on localhost:8989")
	routers.StartServer().Run(":8989")
}
