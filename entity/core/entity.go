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
	Status string
	DetailsBook []DetailsBookCore
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type DetailsBookCore struct {
	ID        string
	Date      string
	Time      string
	BookingID string
	ServiceID string
	CreatedAt time.Time
	UpdatedAt time.Time
}
