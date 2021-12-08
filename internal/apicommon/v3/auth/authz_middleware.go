package auth

import (
	"fmt"

	"github.com/labstack/echo/v4"
	v3 "github.com/cloudfoundry/go-cf-api/internal/apicommon/v3"
)

//nolint:gochecknoglobals
var (
	readScopes  = []Scope{Read, Admin, AdminReadOnly, GlobalAuditor}
	writeScopes = []Scope{Write, Admin}
)

func NewRequiresReadMiddleware() echo.MiddlewareFunc {
	return newRequiresScopesMiddleware(readScopes, "read")
}

func NewRequiresWriteMiddleware() echo.MiddlewareFunc {
	return newRequiresScopesMiddleware(writeScopes, "write")
}

func newRequiresScopesMiddleware(scopes []Scope, verb string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			for _, scope := range scopes {
				if hasScope(scope, ctx) {
					return next(ctx)
				}
			}
			return v3.NotAuthorized(fmt.Errorf("user lacks sufficient scope to %s resource", verb))
		}
	}
}
