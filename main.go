package main

import (
	"basic-trade-api/controllers"
	"basic-trade-api/database"
	"basic-trade-api/routers"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	db := database.DBConnection()
	controllers.EstablishConnection(controllers.NewConnection(db))
	// rand seed for generating random string
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Listening on localhost:8989")
	routers.StartServer().Run(":8989")
}
