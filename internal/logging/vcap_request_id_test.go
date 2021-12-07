//go:build unit
// +build unit

package logging_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/logging"
)

func TestVcapRequestIdSuite(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		externalHeaders map[string]string
		expectedLength  int
		expectedPrefix  string
	}{
		"no external header": {
			expectedLength: 36,
		},
		"external vcap request id": {
			externalHeaders: map[string]string{logging.HeaderVcapRequestID: "123-abc"}, expectedLength: 45, expectedPrefix: "123-abc::",
		},
		"external request id": {
			externalHeaders: map[string]string{echo.HeaderXRequestID: "123-abc"}, expectedLength: 45, expectedPrefix: "123-abc::",
		},
		"non-numeric external id": {
			externalHeaders: map[string]string{logging.HeaderVcapRequestID: "123-abc.?:"}, expectedLength: 45, expectedPrefix: "123-abc::",
		},
		"over 255 char external id": {
			externalHeaders: map[string]string{logging.HeaderVcapRequestID: strings.Repeat("x", 500)},
			expectedLength:  255 + 2 + 36,
			expectedPrefix:  strings.Repeat("x", 255) + "::",
		},
	}

	for testCaseName, testCase := range cases {
		t.Run(testCaseName, func(t *testing.T) {
			e := echo.New()
			request := httptest.NewRequest(http.MethodGet, "/something", nil)
			recorder := httptest.NewRecorder()
			context := e.NewContext(request, recorder)
			handler := func(c echo.Context) error {
				return c.String(http.StatusOK, "")
			}
			for headerName, headerValue := range testCase.externalHeaders {
				request.Header.Set(headerName, headerValue)
			}
			err := logging.NewVcapRequestID()(handler)(context)
			require.NoError(t, err)
			vcapRequestID := recorder.Header().Get(logging.HeaderVcapRequestID)
			require.Equal(t, testCase.expectedLength, len(vcapRequestID))
			require.True(t, strings.HasPrefix(vcapRequestID, testCase.expectedPrefix))
		})
	}
}
