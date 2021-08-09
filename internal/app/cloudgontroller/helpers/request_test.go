package helpers_test

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/helpers"
)

func TestNormalizeAuthScheme(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		authHeader           string
		normalizedAuthHeader string
	}{
		"Auth scheme beginning with upper case character": {
			authHeader:           "Bearer 123",
			normalizedAuthHeader: "bearer 123",
		},
		"Auth scheme beginning with lower case character": {
			authHeader:           "bearer 123",
			normalizedAuthHeader: "bearer 123",
		},
		"No auth scheme": {
			authHeader:           "123",
			normalizedAuthHeader: "123",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest("", "/", nil)
			req.Header.Set(echo.HeaderAuthorization, tc.authHeader)
			c := echo.New().NewContext(req, nil)

			helpers.NormalizeAuthScheme(c)

			assert.Equal(t, tc.normalizedAuthHeader, c.Request().Header.Get(echo.HeaderAuthorization))
		})
	}
}
