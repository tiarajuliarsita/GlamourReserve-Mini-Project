package entity

import (
	"time"
)

type UserCore struct {
	ID      string
	UserName  string
	Email     string
	Password  string
	Phone string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
