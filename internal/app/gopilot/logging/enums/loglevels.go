package enums

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"go.uber.org/zap/zapcore"
)

type levels struct {
	Debug  int8
	Info   int8
	Warn   int8
	Error  int8
	DPanic int8
	Panic  int8
	Fatal  int8
}

var Levels = &levels{
	Debug:  -1,
	Info:   0,
	Warn:   1,
	Error:  2,
	DPanic: 3,
	Panic:  4,
	Fatal:  5,
}

func StringToLogLevel(level string) (zapcore.Level, error) {
	val := reflect.Indirect(reflect.ValueOf(Levels))
	id, err := strconv.Atoi(fmt.Sprintf("%v", val.FieldByName(strings.Title(strings.ToLower(string(level))))))
	if err != nil {
		return zapcore.Level(id), fmt.Errorf("LogLevel %v is not a valid level", level)
	} else {
		return zapcore.Level(id), nil
	}
}
