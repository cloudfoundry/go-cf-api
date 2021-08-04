package v3

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/swagger"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/controllers"
)

func RegisterHealthHandler(e *echo.Echo) {
	// Health
	e.GET("healthz", controllers.GetHealth)
}

func RegisterV3Handlers(prefix string, e *echo.Echo) {
	var uaakey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtDbj14vkov7AFq6fVEJN
+1qqwfgGDIj7k3sohtwbWvTHQA3kWXb6ffUyVtUrYWk5LcQLT2v4fEzNInCif082
yGwKdtTRG8kXDEyXT1CyydK+LznneH+AOxno3I8QKxm3Vhl1pcV0ZR9p928FdsHn
wFDy2hnPFLmEnKn5HEU+zzE39NgO/h8ZJvcA2sO+PgtJd8YY2gXWiHzzrbCH9+Kd
rJprXeFqA8WIRWvBYckzXZFQtwqk9blX0CndgNGog7dwDmoLYdHa6e5p5tSuqhAp
QRonhJgGhX1Kj94TB+msTsucDExLPZMTx/OKNVyTFwsOmTrQPl0Fb1pXMf5S0d7E
tQIDAQAB
-----END PUBLIC KEY-----
`
	var publicKey, _ = jwt.ParseRSAPublicKeyFromPEM([]byte(uaakey))
	// Restricted group
	r := e.Group(prefix)
	{
		config := middleware.JWTConfig{
			SigningKey: publicKey,
			SigningMethod: "RS256",
		}
		r.Use(middleware.JWTWithConfig(config))

		// Buildpacks
		r.GET(fmt.Sprintf("/buildpacks"), controllers.GetBuildpacks)
		r.GET(fmt.Sprintf("/buildpacks/:guid"), controllers.GetBuildpack)
	}
}

func RegisterV3DocumentationHandlers(prefix string, e *echo.Echo) {
	// Serve Swagger-UI API Docs
	e.GET(prefix, func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s/index.html", prefix))
	})
	e.GET(fmt.Sprintf("%s/", prefix), func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s/index.html", prefix))
	})
	e.GET(fmt.Sprintf("%s/*", prefix), echoSwagger.WrapHandler)
}
