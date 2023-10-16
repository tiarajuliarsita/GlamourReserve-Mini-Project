package request

type VariantRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	ServiceID   string `json:"service_id" form:"service_id"`
}

type ServiceRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	// Image       *multipart.FileHeader `form:"image"`
}

type UserRequest struct {
	UserName string `json:"user_name" form:"user_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
}

