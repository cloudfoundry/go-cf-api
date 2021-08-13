package logging_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
)

func TestHeadersFieledValues(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("remote_ip", logging.RemoteIPField)
	assert.Equal("time", logging.TimeField)
	assert.Equal("host", logging.HostField)
	assert.Equal("request", logging.RequestField)
	assert.Equal("user_agent", logging.UserAgentField)
	assert.Equal("status", logging.StatusField)
	assert.Equal("size", logging.SizeField)
	assert.Equal("request_id", logging.RequestIDField)
}
