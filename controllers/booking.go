package controllers

import (
	"booking-fields-app/models"
	"time"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type BookingRequest struct {
	FieldID     uint      `json:"field_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}

func CreateBooking(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req BookingRequest
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
		}

		if req.EndTime.Before(req.StartTime) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "End time must be after start time"})
		}

		var field models.Field
		if err := db.First(&field, req.FieldID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Field not found"})
		}

		var conflict models.Booking
		err := db.Where("field_id = ? AND NOT (end_time <= ? OR start_time >= ?)", req.FieldID, req.StartTime, req.EndTime).First(&conflict).Error
		if err == nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Field is already booked at that time"})
		}
		if err != nil && err != gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not check for booking conflicts: " + err.Error()})
		}

		userID := c.Locals("user_id").(uint)
		booking := models.Booking{
			UserID:    userID,
			FieldID:   req.FieldID,
			StartTime: req.StartTime,
			EndTime:   req.EndTime,
			Status: "pending",
		}
		if err := db.Create(&booking).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create booking: " + err.Error()})
		}
		if err := db.Preload("Field").Preload("User").First(&booking, booking.ID).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve booking details: " + err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(booking)
	}
}

func ListBookings(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var bookings []models.Booking
		userID := c.Locals("user_id").(uint)
		if err := db.Where("user_id = ?", userID).Preload("Field").Preload("User").Find(&bookings).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve bookings: " + err.Error()})
		}
		return c.JSON(bookings)
	}
}