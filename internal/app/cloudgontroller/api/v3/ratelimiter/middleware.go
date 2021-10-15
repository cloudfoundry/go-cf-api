package ratelimiter

import (
	"github.com/labstack/echo/v4"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/auth"
)

func NewRateLimiterMiddleware(generalRateLimiter, unauthenticatedRateLimiter RateLimiter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if auth.IsAdmin(ctx) || auth.IsAdminReadOnly(ctx) {
				return next(ctx)
			}
			limiter := generalRateLimiter
			identifier, ok := ctx.Get(auth.Username).(string)
			if !ok {
				identifier = ctx.RealIP()
				limiter = unauthenticatedRateLimiter
			}
			allowed, headers, err := limiter.Check(identifier)
			if err != nil {
				return err
			}
			for header, headerValue := range headers {
				ctx.Response().Header().Set(header, headerValue)
			}

			if !allowed {
				ctx.Response().Header().Set("Retry-After", ctx.Response().Header().Get("X-RateLimit-Reset"))
				return v3.TooManyRequests(err)
			}
			limiter.Increment(identifier)
			return next(ctx)
		}
	}
}
