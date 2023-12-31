package request

type ServiceRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	// Image       string `form:"image"`
}

type UserRequest struct {
	UserName string `json:"user_name" form:"user_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
}

type BookDataRequest struct {
	Service_id   string `json:"service_id"`
	ServiceStart string `json:"service_start_time"`
	ServiceEnd   string `json:"service_end_time"`
}

type BookingRequest struct {
	Details []BookDataRequest `json:"details"`
}

type NewStatusReq struct {
	Status string `json:"new_status"`
	NoInvoice string `json:"no_invoice"`
}

type AskBeautyReq struct {
	Question    string  `json:"Question"`
}
