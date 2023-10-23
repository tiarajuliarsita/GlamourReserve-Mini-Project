package response

import (
	"time"
)

type UserRespon struct {
	ID        string    `json:"id"`
	UserName  string    `json:"user_name" `
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ServiceRespon struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
	Price       int    `json:"price"`
	// Image       string    `json:"image"`
	// Variants  []models.Variant `json:"variants" `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookingRespon struct {
	ID             string `json:"id"`
	Name           string `json:"name,omitempty"`
	Status         string
	InvoiceNumb    string                `json:"invoice_num"`
	Total          int                   `json:"total"`
	DetailsBooking []DetailBookingRespon `json:"details,omitempy"`
	CreatedAt      time.Time             `json:"created_at"`
}

type DetailBookingRespon struct {
	ID           string    `json:"id"`
	ServiceID    string    `json:"service_id"`
	Name         string    `json:"name"`
	Price        int       `json:"price"`
	ServiceStart string `json:"service_start_time"`
	ServiceEnd   string `json:"service_end_time"`
}
