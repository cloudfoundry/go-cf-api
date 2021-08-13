package logging

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger is a middleware and zap to provide an "access log" like logging for each request.
func NewEchoZapLogger(log *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			req := c.Request()
			res := c.Response()
			vcapRequestID := res.Header().Get(HeaderVcapRequestID)
			fields := []zapcore.Field{
				zap.String(RemoteIPField, c.RealIP()),
				zap.String(TimeField, time.Since(start).String()),
				zap.String(HostField, req.Host),
				zap.String(RequestField, fmt.Sprintf("%s %s", req.Method, req.RequestURI)),
				zap.String(UserAgentField, req.UserAgent()),
				zap.String(RequestIDField, vcapRequestID),
			}

			log.Debug("Request received", fields...)

			err := next(c)
			if err != nil {
				log = log.With(zap.Error(err))
				c.Error(err)
			}

			fields = []zapcore.Field{
				zap.String(RemoteIPField, c.RealIP()),
				zap.String(TimeField, time.Since(start).String()),
				zap.String(HostField, req.Host),
				zap.String(RequestField, fmt.Sprintf("%s %s", req.Method, req.RequestURI)),
				zap.Int(StatusField, res.Status),
				zap.Int64(SizeField, res.Size),
				zap.String(UserAgentField, req.UserAgent()),
				zap.String(RequestIDField, vcapRequestID),
			}

			n := res.Status
			switch {
			case n >= 500: //nolint:gomnd // HTTP error code ranges are well understood
				log.Error("Server error", fields...)
			case n >= 400: //nolint:gomnd // HTTP error code ranges are well understood
				log.Warn("Client error", fields...)
			case n >= 300: //nolint:gomnd // HTTP error code ranges are well understood
				log.Info("Redirection", fields...)
			default:
				log.Info("Success", fields...)
			}

			return nil
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
