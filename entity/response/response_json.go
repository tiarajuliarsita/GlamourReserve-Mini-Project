package response

import "github.com/labstack/echo/v4"

func RespondJSON(e echo.Context, status int, message string, data interface{}) error {
	response := echo.Map{
		"status": message,
		"data":   data,
	}
	return e.JSON(status, response)
}

func RespJSONWithUserName(e echo.Context, status int, message string, data interface{}, username string) error {
	response := echo.Map{
		"status":    message,
		"user_name": username,
		"data":      data,
	}
	return e.JSON(status, response)
}

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func ResponseJSON(status string, data interface{}) Response {
	return Response{
		Status: status,
		Data:   data,
	}
}
