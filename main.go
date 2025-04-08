package main

import (
	"log"
	"os"

	"finance-tracker/configs"
	"finance-tracker/internal/domain"
	"finance-tracker/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file")
	}

	// Initialize database connection
	db := configs.InitDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Auto migrate schema
	db.AutoMigrate(&domain.User{}, &domain.Transaction{})

	// Initialize router
	r := router.NewRouter(db)

	// Get port from environment variables or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
