package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetHealth godoc
// @Summary Health get health
// @Description Health endpoint to check platform availability
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 500 {object} HTTPError
// @Router /healthz [get]
func GetHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}