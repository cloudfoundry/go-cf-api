package v1

import (
	"fmt"
	"net/http"

	"github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/api/v1/controllers"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterHandlers(prefix string, e *echo.Echo) {
	// Pilots
	e.GET(fmt.Sprintf("%s/pilots", prefix), controllers.ListPilots)
	e.GET(fmt.Sprintf("%s/pilot", prefix), controllers.InsertPilot)
}

func RegisterDocumentationHandlers(prefix string, e *echo.Echo) {
	// Serve Swagger-UI API Docs
	e.GET(prefix, func(c echo.Context) error {
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s/index.html", prefix))
		return nil
	})
	e.GET(fmt.Sprintf("%s/", prefix), func(c echo.Context) error {
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s/index.html", prefix))
		return nil
	})
	e.GET(fmt.Sprintf("%s/*", prefix), echoSwagger.WrapHandler)
}
