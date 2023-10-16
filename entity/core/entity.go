package core

import (
	"glamour_reserve/entity/models"
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
	Variants    []models.Variant
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type VariantCore struct {
	ID          string
	Name        string
	Description string
	Price       int
	ServiceID   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
