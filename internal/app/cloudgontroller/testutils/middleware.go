// +build unit

package testutils

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type HandlerFunc struct {
	mock.Mock
}

func (m *HandlerFunc) Next(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}
