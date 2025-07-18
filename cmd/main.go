package main

import (
	"github.com/Mukam21/REST-API_online-marketplace/pkg/database"
	"github.com/Mukam21/REST-API_online-marketplace/pkg/router"
)

func main() {
	db := database.InitDB()

	r := router.SetupRouter(db)

	r.Run(":8080")
}
