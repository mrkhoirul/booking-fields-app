package routes

import (
	"booking-fields-app/controllers"
	"booking-fields-app/middleware"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	// Public routes
	app.Post("/login", controllers.Login(db))
	app.Post("/register", controllers.Register(db))

	// Field routes
	fieldGroup := app.Group("/fields")
	fieldGroup.Get("/", controllers.ListFields(db))
	fieldGroup.Get("/:id", controllers.GetField(db))
	// Protected admin routes
	fieldGroup.Post("/", middleware.JWTProtected(), middleware.AdminOnly(), controllers.CreateField(db))
	fieldGroup.Put("/:id", middleware.JWTProtected(), middleware.AdminOnly(), controllers.UpdateField(db))
	fieldGroup.Delete("/:id", middleware.JWTProtected(), middleware.AdminOnly(), controllers.DeleteField(db))

	// Booking routes
	bookingGroup := app.Group("/bookings", middleware.JWTProtected())
	bookingGroup.Post("/", controllers.CreateBooking(db))
	bookingGroup.Get("/", controllers.ListBookings(db))

	// Payment routes
	paymentGroup := app.Group("/payments", middleware.JWTProtected())
	paymentGroup.Post("/", controllers.CreatePayment(db))
}
