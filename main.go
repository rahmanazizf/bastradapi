package main

import (
	"basic-trade-api/database"
	"basic-trade-api/routers"
	"fmt"
)

func main() {
	// db := database.DBConnection()
	database.DBConnection()
	// controllers.EstablishConnection(controllers.NewConnection(db))
	// var connection *models.Connection
	// connection.ConnectToDB(db)
	// rand seed for generating random string
	// rand.Seed(time.Now().UnixNano())
	fmt.Println("Listening on localhost:8989")
	routers.StartServer().Run(":8989")
}
