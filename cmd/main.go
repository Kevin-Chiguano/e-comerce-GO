package main

import (
	"ecommerce-manager/internal/database"
	"ecommerce-manager/internal/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Println("No .env file found at configs/.env")
		if err = godotenv.Load(); err != nil {
			log.Println("No .env file found in current directory")
		}
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Static files
	r.Static("/static", "./web")

	// Routes
	routes.SetupRoutes(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on http://localhost:%s", port)
	r.Run(":" + port)
}