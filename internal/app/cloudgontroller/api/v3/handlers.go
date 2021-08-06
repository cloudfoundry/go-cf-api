package v3

import (
	"context"
	"errors"
	"fmt"
	jwtv3 "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lestrrat-go/jwx/jwk"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/controllers"
)

var cloudGontrollerConfig *config.CloudgontrollerConfig

func RegisterHealthHandler(e *echo.Echo) {
	// Health
	e.GET("healthz", controllers.GetHealth)
}

func RegisterV3Handlers(prefix string, e *echo.Echo, conf *config.CloudgontrollerConfig) {
	// Restricted group
	cloudGontrollerConfig = conf
	restrictedGroup := e.Group(prefix)
	{
		config := middleware.JWTConfig{
			SigningMethod: "RS256",
			KeyFunc: getUaaKey,
		}
		restrictedGroup.Use(middleware.JWTWithConfig(config))

		// Buildpacks
		restrictedGroup.GET("/buildpacks", controllers.GetBuildpacks)
		restrictedGroup.GET("/buildpacks/:guid", controllers.GetBuildpack)
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

func getUaaKey(token *jwtv3.Token) (interface{}, error) {
	keySet, err := jwk.Fetch(context.Background(), fmt.Sprintf("%s/token_keys", cloudGontrollerConfig.Uaa.Url))
	if err != nil {
		return nil, err
	}
	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have a key ID in the kid field")
	}
	key, found := keySet.LookupKeyID(keyID)
	if !found {
		return nil, fmt.Errorf("unable to find key %q", keyID)
	}
	var pubkey interface{}
	if err := key.Raw(&pubkey); err != nil {
		return nil, fmt.Errorf("Unable to get the public key. Error: %s", err.Error())
	}
	return pubkey, nil
}