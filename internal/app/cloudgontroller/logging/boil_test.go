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

func TestBoilerForBothCases(t *testing.T) {
	t.Parallel()
	// with parameters redacted
	testBoilLoggerParams(t, true)
	// without parameters redacted
	testBoilLoggerParams(t, false)
}

func testBoilLoggerParams(t *testing.T, redactParams bool) {
	t.Helper()
	obs, logs := observer.New(zap.DebugLevel)
	assertEmpty(t, logs)
	assert.NoError(t, obs.Sync(), "Unexpected failure in no-op Sync")
	logger := zap.New(obs)

	bl := logging.NewBoilLogger(redactParams, logger)
	_, err := bl.Write([]byte("SELECT 1 from where you would like"))
	assert.NoError(t, err)
	_, err = bl.Write([]byte(""))
	assert.NoError(t, err)

	params := ""
	if redactParams {
		params = "REDACTED"
	}
	want := observer.LoggedEntry{
		Entry: zapcore.Entry{Level: zap.DebugLevel, Message: "SQL LOG"},
		Context: []zapcore.Field{
			zap.String("query", "SELECT 1 from where you would like"),
			zap.String("params", params),
		},
	}
	assert.Len(t, logs.AllUntimed(), 1)
	log := logs.AllUntimed()[0]
	assertEqual(t, want, log)
}

func assertEqual(t *testing.T, want, log observer.LoggedEntry) {
	t.Helper()
	assert.Equal(t, want.Context, log.Context)
	assert.Equal(t, want.Entry.Level, log.Level)
	assert.Equal(t, want.Entry.Message, log.Message)
}

//nolint
func assertEmpty(t testing.TB, logs *observer.ObservedLogs) {
	assert.Equal(t, 0, logs.Len(), "Expected empty ObservedLogs to have zero length.")
	assert.Equal(t, []observer.LoggedEntry{}, logs.All(), "Unexpected LoggedEntries in empty ObservedLogs.")
}
