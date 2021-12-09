package logging

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const (
	HeaderVcapRequestID = "X-Vcap-Request-Id"
	MaxSizeLimit        = 255
)

var vcapRequestIDBadCharRegexp = regexp.MustCompile(`[^\w-]`) // shorthand for [a-zA-Z0-9_]

func NewVcapRequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			existingRequestID := ctx.Request().Header.Get(HeaderVcapRequestID)
			if len(existingRequestID) == 0 {
				existingRequestID = ctx.Request().Header.Get(echo.HeaderXRequestID)
			}
			requestID := buildRequestID(existingRequestID)

			ctx.Response().Header().Set(HeaderVcapRequestID, requestID)
			return next(ctx)
		}
	}
}

func buildRequestID(vcapRequestID string) string {
	guid := uuid.New().String()
	if len(vcapRequestID) == 0 {
		vcapRequestID = guid
		return vcapRequestID
	}

	// append guid to first guid of existing request id
	vcapRequestID = vcapRequestIDBadCharRegexp.ReplaceAllString(vcapRequestID, "")
	if len(vcapRequestID) > MaxSizeLimit {
		vcapRequestID = vcapRequestID[:MaxSizeLimit]
	}
	vcapRequestID = fmt.Sprintf("%s::%s", vcapRequestID, guid)
	return vcapRequestID
}
