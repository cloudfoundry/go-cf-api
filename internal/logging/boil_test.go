//go:build unit
// +build unit

package logging_test

import (
	"testing"

	"github.com/cloudfoundry/go-cf-api/internal/logging"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

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

	for testCaseName, testCase := range tests {
		t.Run(testCaseName, func(t *testing.T) {
			obs, logs := observer.New(zap.DebugLevel)
			assert.Empty(t, logs)
			logger := zap.New(obs)
			bl := logging.NewBoilLogger(testCase.redactParams, logger)

			for _, call := range testCase.writeCalls {
				_, err := bl.Write([]byte(call))
				require.NoError(t, err)
			}

			assert.Len(t, logs.AllUntimed(), len(testCase.expectedEntries))
			for idx, log := range logs.AllUntimed() {
				assert.True(t, log.Caller.Defined)
				assert.Equal(t, testCase.expectedEntries[idx].Level, log.Level)
				assert.Equal(t, testCase.expectedEntries[idx].Message, log.Message)
				assert.Equal(t, testCase.expectedEntries[idx].Context, log.Context)
			}
		})
	}
}
