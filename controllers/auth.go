package controllers

import (
	"booking-fields-app/models"
	"booking-fields-app/utils"
	"os"
	"strconv"
	"time"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req RegisterRequest
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}
		if req.Email == "" || req.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email and password are required"})
		}
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
		}
		role := req.Role
		if role == "" {
			role = "user"
		}
		user := models.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: hashedPassword,
			Role:     role,
		}
		if err := db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user: " + err.Error()})
		}
		user.Password = ""
		return c.Status(fiber.StatusCreated).JSON(user)
	}
}

func Login(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req LoginRequest
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}
		var user models.User
		if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
		}
		if err := utils.CheckPasswordHash(req.Password, user.Password); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
		}

		secret := os.Getenv("JWT_SECRET")
		expirationHoursStr := os.Getenv("JWT_EXPIRE_HOURS")
		expirationHours := 72
		if expirationHoursStr != "" {
			if hrs, err := strconv.Atoi(expirationHoursStr); err == nil {
				expirationHours = hrs
			}
		}
		claims := jwt.MapClaims{
			"user_id": user.ID,
			"role":    user.Role,
			"exp":     time.Now().Add(time.Duration(expirationHours) * time.Hour).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, err := token.SignedString([]byte(secret))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
		}
		return c.JSON(fiber.Map{"token": signedToken})
	}
}