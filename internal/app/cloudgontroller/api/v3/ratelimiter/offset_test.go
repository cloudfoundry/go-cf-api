// +build unit

package ratelimiter_test

import (
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	. "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/ratelimiter"
)

func TestUserHashOffsetProducesOffsetsWithinInterval(t *testing.T) {
	interval := time.Hour
	for i := 0; i < 1000; i++ {
		guid, err := uuid.NewV4()
		assert.NoError(t, err)
		offset := UserHashOffset(guid.String(), interval)
		now := time.Now()
		assert.WithinDuration(t, now, now.Add(offset), interval)
	}
}

func TestUserHashOffsetProducesSameOffsetForSameInputs(t *testing.T) {
	offset1 := UserHashOffset("user-guid", time.Hour)
	offset2 := UserHashOffset("user-guid", time.Hour)
	assert.Equal(t, offset1, offset2)
}
