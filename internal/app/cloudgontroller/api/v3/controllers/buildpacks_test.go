// +build unit

package controllers_test

import (

	// models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetBuildpack(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/buildpacks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:guid")
	c.SetParamNames("guid")
	c.SetParamValues("aeb7c510-2bec-4508-8309-fda7b4a67a1d")

	// Test logic missing

	assert.True(t, true)
}
