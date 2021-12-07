package auth

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	Username = "username"
	Scopes   = "scopes"
)

type CFClaims struct {
	Audience  []string `json:"aud"`
	ClientID  string   `json:"client_id,omitempty"`
	ExpiresAt int64    `json:"exp,omitempty"`
	ID        string   `json:"jti,omitempty"`
	IssuedAt  int64    `json:"iat,omitempty"`
	Issuer    string   `json:"iss,omitempty"`
	Subject   string   `json:"sub,omitempty"`
	UserID    string   `json:"user_id"`
	Username  string   `json:"user_name"`
	Scopes    []string `json:"scope"`
}

func (c *CFClaims) Valid() error {
	// We would like to use jwt.StandardClaims but that does not accept a []string Audience
	// so we build an instance with the same variables and validate it
	standard := jwt.StandardClaims{
		ExpiresAt: c.ExpiresAt,
		Id:        c.ID,
		IssuedAt:  c.IssuedAt,
		Issuer:    c.Issuer,
		Subject:   c.Subject,
	}

	return standard.Valid()
}

func NewJWTMiddleware(keyFunc jwt.Keyfunc) echo.MiddlewareFunc {
	authConfig := middleware.JWTConfig{
		SigningMethod:  "RS256",
		KeyFunc:        keyFunc,
		Claims:         &CFClaims{},
		SuccessHandler: NewSuccessHandler(),
	}
	return middleware.JWTWithConfig(authConfig)
}

func NewSuccessHandler() middleware.JWTSuccessHandler {
	return func(ctx echo.Context) {
		user, ok := ctx.Get("user").(*jwt.Token)
		if !ok {
			ctx.Error(errors.New("couldn't get user from context"))
		}
		claims, ok := user.Claims.(*CFClaims)
		if !ok {
			ctx.Error(errors.New("couldn't get user claims from context"))
		}

		ctx.Set(Username, claims.UserID)
		ctx.Set(Scopes, claims.Scopes)
	}
}
