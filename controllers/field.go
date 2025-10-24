package controllers

import (
	"booking-fields-app/models"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func CreateField(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		field := new(models.Field)
		if err := c.Bind().Body(field); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
		}
		if err := db.Create(&field).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create field: " + err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(field)
	}
}

func ListFields(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var fields []models.Field
		if err := db.Find(&fields).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve fields: " + err.Error()})
		}
		return c.JSON(fields)
	}
}

func GetField(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Params("id")
		var field models.Field
		if err := db.First(&field, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Field not found"})
		}
		return c.JSON(field)
	}
}

func UpdateField(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Params("id")
		var field models.Field
		if err := db.First(&field, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Field not found"})
		}
		var input models.Field
		if err := c.Bind().Body(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
		}
		field.Name = input.Name
		field.PricePerHour = input.PricePerHour
		field.Location = input.Location
		if err := db.Save(&field).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update field: " + err.Error()})
		}
		return c.JSON(field)
	}
}

func DeleteField(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Params("id")
		var field models.Field
		if err := db.First(&field, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Field not found"})
		}
		if err := db.Delete(&models.Field{}, id).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete field: " + err.Error()})
		}
		return c.JSON(fiber.Map{"message": "Field deleted successfully"})
	}
}