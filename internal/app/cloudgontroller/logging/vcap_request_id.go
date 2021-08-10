package logging

import (
	"fmt"
	"regexp"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const (
	HeaderVcapRequestID = "X-Vcap-Request-Id"
	MaxSizeLimit        = 255
)

var vcapRequestIDBadCharRegexp = regexp.MustCompile(`[^\w-]`) // shorthand for [a-zA-Z0-9_]

func NewVcapRequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			existingRequestID := c.Request().Header.Get(HeaderVcapRequestID)
			if len(existingRequestID) == 0 {
				existingRequestID = c.Request().Header.Get(echo.HeaderXRequestID)
			}
			requestID, err := buildRequestID(existingRequestID)
			if err != nil {
				return err
			}
			c.Response().Header().Set(HeaderVcapRequestID, requestID)
			return next(c)
		}
	}
}

func buildRequestID(vcapRequestID string) (string, error) {
	guid, err := uuid.NewV4()
	if err != nil {
		zap.L().Error("guid generation failed", zap.Error(err))
		return "", err
	}
	guidStr := guid.String()

	if len(vcapRequestID) == 0 {
		vcapRequestID = guidStr
		return vcapRequestID, nil
	}

	// append guid to first guid of existing request id
	vcapRequestID = vcapRequestIDBadCharRegexp.ReplaceAllString(vcapRequestID, "")
	if len(vcapRequestID) > MaxSizeLimit {
		vcapRequestID = vcapRequestID[:MaxSizeLimit]
	}
	vcapRequestID = fmt.Sprintf("%s::%s", vcapRequestID, guidStr)
	return vcapRequestID, nil
}
