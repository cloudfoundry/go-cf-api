// +build unit

package timefilters //nolint:testpackage // we have to assign package level vars due to models using static functions

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
		"equal": {
			query:              fmt.Sprintf("created_ats=%s&updated_ats=%s", now, now),
			expectedCreatedAts: []TimeFilter{{Operator: "", Timestamp: nowT}},
			expectedUpdatedAts: []TimeFilter{{Operator: "", Timestamp: nowT}},
		},
		"less than": {
			query:              fmt.Sprintf("created_ats[lt]=%s&updated_ats[lte]=%s", now, now),
			expectedCreatedAts: []TimeFilter{{Operator: "[lt]", Timestamp: nowT}},
			expectedUpdatedAts: []TimeFilter{{Operator: "[lte]", Timestamp: nowT}},
		},
		"greater than": {
			query:              fmt.Sprintf("created_ats[gt]=%s&updated_ats[gte]=%s", now, now),
			expectedCreatedAts: []TimeFilter{{Operator: "[gt]", Timestamp: nowT}},
			expectedUpdatedAts: []TimeFilter{{Operator: "[gte]", Timestamp: nowT}},
		},
		"repeated field with different comparisons": {
			query: fmt.Sprintf("created_ats[gt]=%s&created_ats[lte]=%s", now, now),
			expectedCreatedAts: []TimeFilter{
				{Operator: "[gt]", Timestamp: nowT},
				{Operator: "[lte]", Timestamp: nowT},
			},
		},
		"unrecognized comparison": {
			query:         fmt.Sprintf("created_ats[foo]=%s", now),
			expectedError: fmt.Errorf("unrecognized comparison operator '[foo]'"),
		},
		"malformed filter": {
			query:         fmt.Sprintf("created_ats[gte=%s", now),
			expectedError: fmt.Errorf("could not parse 'created_ats' filter created_ats[gte"),
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
