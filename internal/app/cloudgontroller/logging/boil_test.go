// +build unit

package logging_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestBoilLogger(t *testing.T) {
	t.Parallel()

	obs, logs := observer.New(zap.DebugLevel)
	assert.NoError(t, obs.Sync(), "Unexpected failure in no-op Sync")

	logger := zap.New(obs)
	bl := logging.NewBoilLogger(true, logger)
	_, err := bl.Write([]byte("SELECT * FROM \"buildpacks\" LIMIT 50;"))
	assert.NoError(t, err)
	_, err = bl.Write([]byte(""))
	assert.NoError(t, err)

	want := observer.LoggedEntry{
		Entry: zapcore.Entry{Level: zap.DebugLevel, Message: "SQL LOG"},
		Context: []zapcore.Field{
			zap.String("query", "SELECT * FROM \"buildpacks\" LIMIT 50;"),
			zap.String("params", "REDACTED"),
		},
	}
	assert.Len(t, logs.AllUntimed(), 1)
	log := logs.AllUntimed()[0]
	assert.Equal(t, want.Context, log.Context)
	assert.Equal(t, want.Entry.Level, log.Level)
	assert.Equal(t, want.Entry.Message, log.Message)
}
