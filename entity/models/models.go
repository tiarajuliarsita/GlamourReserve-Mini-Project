package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"not null;primary key"`
	UserName  string `gorm:"not null" valid:"required~your user name is required"`
	Email     string `gorm:"not null;unique" valid:"required~your email is required, email~invalid email format"`
	Password  string `gorm:"not null" valid:"required~your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Phone     string `gorm:"not null" valid:"required~your phone is required"`
	Role      string `gorm:"type:ENUM('user', 'admin');not null;default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Service struct {
	ID          string `gorm:"not null;type:varchar(255);primary key" json:"id" form:"id"`
	Name        string `gorm:"not null;unique" valid:"required~your name is required" json:"name" form:"name"`
	Description string `gorm:"not null" valid:"required~your description is required" json:"description" form:"description"`
	Price       int    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Booking struct {
	ID             string `gorm:"not null;primary key; type:varchar(255)"`
	UserID         string `valid:"required~your user id is required"`
	InvoiceNumb    string `gorm:"not null;unique" valid:"required~your invoice is required"`
	Total          int    `gorm:"not null"`
	Status         string `gorm:"type:ENUM('pending', 'done');not null;default:'pending'"`
	DetailsBooking []DetailBooking
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type DetailBooking struct {
	ID           string    `gorm:"not null;primary key"`
	DateTime     time.Time `gorm:"not null;type:time" validate:"required" valid:"required~your date_time is required"`
	TimeExpected time.Time `gorm:"not null;type:time" validate:"required" valid:"required~your time_expected is required"`
	BookingID    string    `gorm:"type:varchar(255)"`
	ServiceID    string    `gorm:"type:varchar(255)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
