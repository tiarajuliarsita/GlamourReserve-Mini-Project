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
	Role      string
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
	Status      string
	DetailsBook []DetailsBookCore
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type DetailsBookCore struct {
	ID string
	// Date      time.Time
	ServiceStart string 
	ServiceEnd   string
	BookingID    string
	ServiceID    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type BookingAll struct {
	Name        string    `json:"name"`
	BookingId   string    `json:"booking_id"`
	InvoiceNumb string    `json:"invoice_numb"`
	Total       int       `json:"total"`
	Status      string    `json:"status"`
	Created_at  time.Time `json:"created_at"`
}
