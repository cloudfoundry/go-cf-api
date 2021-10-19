package logging

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger is a middleware and zap to provide an "access log" like logging for each request.
func NewEchoZapLogger(baseLogger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			res := c.Response()
			vcapRequestID := res.Header().Get(HeaderVcapRequestID)

			logger := baseLogger.With(
				zap.String(RemoteIPField, c.RealIP()),
				zap.String(HostField, c.Request().Host),
				zap.String(RequestField, fmt.Sprintf("%s %s", c.Request().Method, c.Request().RequestURI)),
				zap.String(UserAgentField, c.Request().UserAgent()),
				zap.String(RequestIDField, vcapRequestID),
			)

			logger.Debug("Request received", zap.String(TimeField, time.Since(start).String()))

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			var responseFields []zapcore.Field
			responseFields = append(responseFields,
				zap.String(TimeField, time.Since(start).String()),
				zap.Int(StatusField, res.Status),
				zap.Int64(SizeField, res.Size),
			)

			responseLogger := logger.With(responseFields...)
			n := res.Status
			switch {
			case n >= 500: //nolint:gomnd // HTTP error code ranges are well understood
				responseLogger.Error("Server error")
			case n >= 400: //nolint:gomnd // HTTP error code ranges are well understood
				responseLogger.Warn("Client error")
			case n >= 300: //nolint:gomnd // HTTP error code ranges are well understood
				responseLogger.Info("Redirection")
			default:
				responseLogger.Info("Success")
			}

			return err
		}
	}
}

/*
https://github.com/brpaz/echozap

The MIT License (MIT)

Copyright (c) Bruno Paz oss@brunopaz.net (https://github.com/brpaz)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
