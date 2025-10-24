package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID			uint			`gorm:"primaryKey" json:"id"`
	Name		string		 	`json:"name"`
	Email		string		 	`gorm:"uniqueIndex" json:"email"`
	Password 	string			`json:"-"`
	Role		string			`json:"role"`
	CreatedAt	time.Time		`json:"created_at"`
	UpdatedAt	time.Time		`json:"updated_at"`
	DeletedAt	gorm.DeletedAt	`gorm:"index" json:"-"`
}

type Field struct {
	ID				uint			`gorm:"primaryKey" json:"id"`
	Name			string		 	`json:"name"`
	PricePerHour	int64			`json:"price_per_hour"`
	Location		string		 	`json:"location"`
	CreatedAt		time.Time		`json:"created_at"`
	UpdatedAt		time.Time		`json:"updated_at"`
	DeletedAt		gorm.DeletedAt	`gorm:"index" json:"-"`
}

type Booking struct {
	ID			uint			`gorm:"primaryKey" json:"id"`
	UserID		uint			`json:"user_id"`
	User		User			`gorm:"foreignKey:UserID" json:"user,omitempty"`
	FieldID		uint			`json:"field_id"`
	Field		Field			`gorm:"foreignKey:FieldID" json:"field,omitempty"`
	StartTime	time.Time		`json:"start_time"`
	EndTime		time.Time		`json:"end_time"`
	Status		string			`json:"status"`
	CreatedAt	time.Time		`json:"created_at"`
	UpdatedAt	time.Time		`json:"updated_at"`
	DeletedAt	gorm.DeletedAt	`gorm:"index" json:"-"`
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Field{}, &Booking{})
}