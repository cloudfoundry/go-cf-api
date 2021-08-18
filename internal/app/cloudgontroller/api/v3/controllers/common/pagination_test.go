package common_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	commoncontroller "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/controllers/common"
)

func TestDefaultPaginationValues(t *testing.T) {
	got := commoncontroller.DefaultPagination()
	assert.Equal(t, 1, got.Page)
	assert.Equal(t, uint16(50), got.PerPage)
}

func TestPaginationParamsValidationInBounds(t *testing.T) {
	t.Parallel()
	const MaxInt = int(^uint(0) >> 1) // Depending on Architecture Uint can be 32 or 64 bit so we need to get MaxUint dynamically
	tests := map[string]commoncontroller.PaginationParams{
		"with lowest possible Page and PerPage":  {1, 1},
		"with largest possible Page and PerPage": {MaxInt, 5000},
	}
	for name, testcase := range tests {
		t.Run(name, func(t *testing.T) {
			got := validator.New().Struct(commoncontroller.PaginationParams{Page: testcase.Page, PerPage: testcase.PerPage})
			assert.Nil(t, got, validatorMessage(got))
		})
	}
}

func TestPaginationParamsValidationOutOfBounds(t *testing.T) {
	t.Parallel()
	tests := map[string]commoncontroller.PaginationParams{
		"with Page lower than 1":        {0, 1},
		"with PerPage lower than 1":     {1, 0},
		"with PerPage larger than 5000": {1, 5001},
	}
	for name, testcase := range tests {
		t.Run(name, func(t *testing.T) {
			got := validator.New().Struct(commoncontroller.PaginationParams{Page: testcase.Page, PerPage: testcase.PerPage})
			assert.Error(t, got, validatorMessage(got))
		})
	}
}

func validatorMessage(validatorError error) string {
	if validatorError != nil {
		return validatorError.Error()
	}
	return ""
}