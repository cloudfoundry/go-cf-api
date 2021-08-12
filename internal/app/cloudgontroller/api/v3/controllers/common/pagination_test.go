package common

import (
	"testing"
)


func TestDefaultPagination(t *testing.T) {
	newDefaultPagination := DefaultPagination()
	if newDefaultPagination.Page != 1 {
		t.Errorf("DefaultPagination.Page = %d; should be 1", newDefaultPagination.Page)
	}
	if newDefaultPagination.PerPage != 50 {
		t.Errorf("DefaultPagination.PerPage = %d; should be 50", newDefaultPagination.PerPage)
	}
}

