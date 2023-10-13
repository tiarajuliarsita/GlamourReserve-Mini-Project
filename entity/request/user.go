package request

type UserRequset struct {
    UserName  string `json:"user_name" form:"user_name"`
    Email     string `json:"email" form:"email"`
    Password  string `json:"password" form:"password"`
    Phone string `json:"phone" form:"phone"`
}
