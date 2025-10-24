package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v3"
	"booking-fields-app/config"
	"booking-fields-app/controllers"
	"booking-fields-app/models"
	"booking-fields-app/routes"
)

func initEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file ftround");
	}
}

func main() {
	initEnv()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := models.MigrateDB(db); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	app := fiber.New()
	controllers.SeedAdmin(db)
	routes.Setup(app, db)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server is running on port  %s\n", port)
	log.Fatal(app.Listen(":" + port))
}