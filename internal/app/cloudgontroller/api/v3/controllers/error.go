package controllers

import "github.com/labstack/echo/v4"

// NewError Function to create an error object
func NewError(c echo.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	c.JSON(status, er)
}

// HTTPError Error Object Returned as json on error
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
