package main

import (
	"log"
	"os"

	"github.com/danielopara/restaurant-api/database"
	"github.com/danielopara/restaurant-api/router"
)

func main() {
	db := database.InitDatabase()

	r := router.Router(db)

	port := os.Getenv("PORT")
	if port == ""{
		port = "3000"
	}
	
	log.Printf("server starting on port %s", port)
	log.Fatal(r.Run(":" + port))
}