package main

import (
	"log"
	"os"

	"github.com/GesaXB/LibayGoManagement/config"
	"github.com/GesaXB/LibayGoManagement/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	var db, err = config.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect Database: %v", err)
	}

	routes.SetupRoutes(server, db)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	if err := server.Run(":" + PORT); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
