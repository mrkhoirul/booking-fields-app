package controllers

import (
	"booking-fields-app/models"
	"booking-fields-app/utils"
	"log"
	"os"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminEmail == "" || adminPassword == "" {
		log.Println("Admin credentials not set in environment variables")
		return
	}
	var existingAdmin models.User
	if err := db.Where("email = ?", adminEmail).First(&existingAdmin).Error; err == nil {
		log.Println("Admin user already exists")
		return
	}
	hashedPassword, err := utils.HashPassword(adminPassword)
	if err != nil {
		log.Printf("Failed to hash admin password: %v", err)
		return
	}
	adminUser := models.User{
		Name:	 "Admin",
		Email:    adminEmail,
		Password: hashedPassword,
		Role:     "admin",
	}
	if err := db.Create(&adminUser).Error; err != nil {
		log.Printf("Failed to create admin user: %v", err)
		return
	}
	log.Println("Admin user created successfully")
}