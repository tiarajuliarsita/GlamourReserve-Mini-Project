package core

import (
	"time"
)

type UserCore struct {
	ID        string
	UserName  string
	Email     string
	Password  string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type ServiceCore struct {
	ID   string
	Name string
	// Image       string
	Price       int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookingCore struct {
	ID          string
	UserID      string
	InvoiceNumb string
	Total       int
	DetailsBook []DetailsBookCore
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type DetailsBookCore struct {
	ID        string
	Date      time.Time
	Time      time.Time
	BookingID string
	ServiceID string
	CreatedAt time.Time
	UpdatedAt time.Time
}
