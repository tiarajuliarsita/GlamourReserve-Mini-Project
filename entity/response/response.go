package response

import (
	"time"
)

type UserRespon struct {
	ID        string    `json:"id"`
	UserName  string    `json:"user_name" `
	Email     string    `json:"email"`
	Phone     string    `json:"phone" `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type ServiceRespon struct {
	ID          string `json:"id"`
	Name        string `json:"name" `
	Description string `json:"description"`
	// Image       string    `json:"image"`
	// Variants  []models.Variant `json:"variants" `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ID          string `gorm:"not null;primary key"`
// 	Name        string `gorm:"not null" valid:"required~your name is required"`
// 	Description string `gorm:"not null;unique" valid:"required~your description is required"`
// 	Image       string `gorm:"not null" valid:"required~your image url is required"`
// 	Variants    []Variant
// 	CreatedAt   time.Time
// 	UpdatedAt   time.Time
// 	DeletedAt   gorm.DeletedAt `gorm:"index"`
// }
