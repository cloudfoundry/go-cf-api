// +build unit

package logging_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestZapLogger(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/something", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	}

	obs, logs := observer.New(zap.DebugLevel)

	logger := zap.New(obs)

	err := logging.NewEchoZapLogger(logger)(h)(c)

	assert.Nil(t, err)
	assert.Equal(t, 2, logs.Len())

	logFieldsStart := logs.AllUntimed()[0].ContextMap()
	assert.Equal(t, "example.com", logFieldsStart["host"])
	assert.NotNil(t, logFieldsStart["time"])
	assert.Equal(t, "GET /something", logFieldsStart["request"])
	assert.NotNil(t, logFieldsStart["user_agent"])

	logFieldsEnd := logs.AllUntimed()[1].ContextMap()
	assert.Equal(t, "192.0.2.1", logFieldsEnd["remote_ip"])
	assert.Equal(t, "example.com", logFieldsEnd["host"])
	assert.NotNil(t, logFieldsEnd["time"])
	assert.Equal(t, "GET /something", logFieldsEnd["request"])
	assert.Equal(t, int64(200), logFieldsEnd["status"])
	assert.Equal(t, int64(0), logFieldsEnd["size"])
	assert.NotNil(t, logFieldsEnd["user_agent"])
	assert.NotNil(t, logFieldsEnd["request_id"])
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
