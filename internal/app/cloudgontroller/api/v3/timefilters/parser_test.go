// +build unit

package timefilters //nolint:testpackage // we have to assign package level vars due to sqlboiler using static functions

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestParseTimeFilters(t *testing.T) {
	nowT := time.Now().UTC().Truncate(time.Second)
	now := nowT.Format(time.RFC3339)

	tests := map[string]struct {
		query              string
		expectedCreatedAts []TimeFilter
		expectedUpdatedAts []TimeFilter
		expectedError      error
	}{
		"filter by time": {
			query: fmt.Sprintf("created_ats=%s&updated_ats=%s", now, now),
			expectedCreatedAts: []TimeFilter{
				{
					Operator:  "",
					Timestamp: nowT,
				},
			},
			expectedUpdatedAts: []TimeFilter{
				{
					Operator:  "",
					Timestamp: nowT,
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "http://test", nil)
			rec := httptest.NewRecorder()
			req.URL.RawQuery = tc.query
			c := e.NewContext(req, rec)

			createdAts, updatedAts, err := ParseTimeFilters(c)
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err)
				return
			}

			assert.NoError(t, err)

			assert.Equal(t, tc.expectedCreatedAts, createdAts)
			assert.Equal(t, tc.expectedUpdatedAts, updatedAts)
		})
	}
}
