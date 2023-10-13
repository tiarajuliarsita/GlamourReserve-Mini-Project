package response

import "time"

type UserRespon struct {
	ID        string    `json:"id"`
	UserName  string    `json:"user_name" `
	Email     string    `json:"email"`
	Phone     string    `json:"phone" `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
