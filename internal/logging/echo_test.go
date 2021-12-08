//go:build unit
// +build unit

package logging_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/cloudfoundry/go-cf-api/internal/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestZapLogger(t *testing.T) {
	t.Parallel()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/something", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	}

	obs, logs := observer.New(zap.DebugLevel)

	logger := zap.New(obs)

	err := logging.NewEchoZapLogger(logger)(handler)(ctx)

	assert.Nil(t, err)
	assert.Equal(t, 2, logs.Len())

	logFieldsStart := logs.AllUntimed()[0].ContextMap()
	assert.Equal(t, "example.com", logFieldsStart[logging.HostField])
	assert.NotNil(t, logFieldsStart[logging.TimeField])
	assert.Equal(t, "GET /something", logFieldsStart[logging.RequestField])
	assert.NotNil(t, logFieldsStart[logging.UserAgentField])

	logFieldsEnd := logs.AllUntimed()[1].ContextMap()
	assert.Equal(t, "192.0.2.1", logFieldsEnd[logging.RemoteIPField])
	assert.Equal(t, "example.com", logFieldsEnd[logging.HostField])
	assert.NotNil(t, logFieldsEnd[logging.TimeField])
	assert.Equal(t, "GET /something", logFieldsEnd[logging.RequestField])
	assert.Equal(t, int64(200), logFieldsEnd[logging.StatusField])
	assert.Equal(t, int64(0), logFieldsEnd[logging.SizeField])
	assert.NotNil(t, logFieldsEnd[logging.UserAgentField])
	assert.NotNil(t, logFieldsEnd[logging.RequestIDField])
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
