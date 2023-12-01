package main

import (
	"basic-trade-api/database"
	"basic-trade-api/routers"
	"fmt"
	"os"
)

var PORT = os.Getenv("PORT")

func main() {
	database.DBConnection()
	fmt.Println(fmt.Sprintf("Listening on :%s", PORT))
	routers.StartServer().Run(fmt.Sprintf(":%s", PORT))
}
