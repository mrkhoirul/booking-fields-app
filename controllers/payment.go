package controllers

import (
	"booking-fields-app/models"
	"strconv"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type PaymentRequest struct {
	BookingID uint   `json:"booking_id"`
}

func CreatePayment(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req PaymentRequest
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
		}

		var booking models.Booking
		if err := db.Preload("Field").First(&booking, req.BookingID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Booking not found"})
		}

		if booking.Status != "pending" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Booking is not pending"})
		}

		booking.Status = "paid"
		if err := db.Save(&booking).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update booking status: " + err.Error()})
		}
		return c.JSON(fiber.Map{
			"message": "Payment successful", 
			"booking_id": strconv.Itoa(int(booking.ID)), 
			"status": booking.Status,
		})
	}
}