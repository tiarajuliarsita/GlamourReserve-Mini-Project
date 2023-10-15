package request

type ServiceRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	// Image       *multipart.FileHeader `form:"image"`
}
