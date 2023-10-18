package request

import "time"

type ServiceRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	// Image       *multipart.FileHeader `form:"image"`
}

type UserRequest struct {
	UserName string `json:"user_name" form:"user_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
}

type BookDataRequest struct {
	Service_id string `json:"service_id"`
	Date       time.Time `json:"date"`
	Time       time.Time `json:"time"`
}

type BookingRequest struct {
	Details []BookDataRequest `json:"details"`
}
