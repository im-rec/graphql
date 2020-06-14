package response

import (
	"github.com/labstack/echo"
	"net/http"
)

type Response struct {
	Success bool         `json:"success"`
	Data    interface{}  `json:"data,omitempty"`
	Errors  *Customerror `json:"error,omitempty"`
}

type Customerror struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Render(c echo.Context, data interface{}) error {
	var response = new(Response)

	response.Success = true
	response.Data = data

	return c.JSON(http.StatusOK, response)
}

//func Error(c echo.Context, err error) error {
	//var response = new(Response)
	//
	//code, status, message := errors.GetError(err)
	//
	//response.Success = false
	//response.Errors = &Customerror{
	//	Status:  status,
	//	Message: message,
	//}
	//
	//return c.JSON(code, response)
//}
