// +build unit

package logging_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

//nolint:funlen
func TestBoilLogger(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		level           zapcore.Level
		redactParams    bool
		writeCalls      []string
		expectedEntries []observer.LoggedEntry
	}{
		"with redacted params": {
			level:        zap.DebugLevel,
			redactParams: true,
			writeCalls: []string{
				"SELECT 1 FROM table WHERE guid = ?;",
				"abc-1234",
			},
			expectedEntries: []observer.LoggedEntry{{
				Entry: zapcore.Entry{Level: zap.DebugLevel, Message: "SQL LOG"},
				Context: []zapcore.Field{
					zap.String("query", "SELECT 1 FROM table WHERE guid = ?;"),
					zap.String("params", "REDACTED"),
				},
			}},
		},
		"with unredacted params": {
			level:        zap.DebugLevel,
			redactParams: false,
			writeCalls: []string{
				"SELECT 1 FROM table WHERE guid = ?;",
				"abc-1234",
			},
			expectedEntries: []observer.LoggedEntry{{
				Entry: zapcore.Entry{Level: zap.DebugLevel, Message: "SQL LOG"},
				Context: []zapcore.Field{
					zap.String("query", "SELECT 1 FROM table WHERE guid = ?;"),
					zap.String("params", "abc-1234"),
				},
			}},
		},
		"when multiple queries are logged": {
			level:        zap.DebugLevel,
			redactParams: false,
			writeCalls: []string{
				"SELECT 1 FROM table WHERE guid = ?;",
				"abc-1234",
				"SELECT 1 FROM other_table WHERE name = ?;",
				"name-5678",
			},
			expectedEntries: []observer.LoggedEntry{
				{
					Entry: zapcore.Entry{Level: zap.DebugLevel, Message: "SQL LOG"},
					Context: []zapcore.Field{
						zap.String("query", "SELECT 1 FROM table WHERE guid = ?;"),
						zap.String("params", "abc-1234"),
					},
				},
				{
					Entry: zapcore.Entry{Level: zap.DebugLevel, Message: "SQL LOG"},
					Context: []zapcore.Field{
						zap.String("query", "SELECT 1 FROM other_table WHERE name = ?;"),
						zap.String("params", "name-5678"),
					},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			obs, logs := observer.New(zap.DebugLevel)
			assert.Empty(t, logs)
			logger := zap.New(obs)
			bl := logging.NewBoilLogger(tc.redactParams, logger)

			for _, call := range tc.writeCalls {
				_, err := bl.Write([]byte(call))
				require.NoError(t, err)
			}

			assert.Len(t, logs.AllUntimed(), len(tc.expectedEntries))
			for idx, log := range logs.AllUntimed() {
				assert.True(t, log.Caller.Defined)
				assert.Equal(t, tc.expectedEntries[idx].Level, log.Level)
				assert.Equal(t, tc.expectedEntries[idx].Message, log.Message)
				assert.Equal(t, tc.expectedEntries[idx].Context, log.Context)
			}
		})
	}
}
