package main

import (
	"github.com/haviz000/restapi-gin-hacktiv8/database"
	"github.com/haviz000/restapi-gin-hacktiv8/routers"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	const PORT = ":8080"
	routers.StartServer().Run(PORT)
}
