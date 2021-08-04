package logging

import (
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"regexp"
)

const (
	HeaderVcapRequestId = "X-Vcap-Request-Id"
)

var vcapRequestIdBadCharRegexp = regexp.MustCompile("[^\\w-]")

func NewVcapRequestId() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			existingRequestId := c.Request().Header.Get(HeaderVcapRequestId)
			if len(existingRequestId) == 0 {
				existingRequestId = c.Request().Header.Get(echo.HeaderXRequestID)
			}
			c.Response().Header().Set(HeaderVcapRequestId, buildRequestId(existingRequestId))
			return next(c)
		}
	}
}

func buildRequestId(vcapRequestId string) string {
	guid, err := uuid.NewV4()
	var guidStr string
	if err == nil {
		guidStr = guid.String()
	} else {
		zap.L().Error("guid generation failed", zap.Error(err))
		guidStr = "???"
	}
	if len(vcapRequestId) == 0 {
		vcapRequestId = guidStr
	} else {
		// append guid to first guid of existing request id
		vcapRequestId = vcapRequestIdBadCharRegexp.ReplaceAllString(vcapRequestId, "")
		if len(vcapRequestId) > 255 {
			vcapRequestId = vcapRequestId[:255]
		}
		vcapRequestId = vcapRequestId + "::" + guidStr
	}
	return vcapRequestId
}
