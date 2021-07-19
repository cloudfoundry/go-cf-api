package helpers

import "go.uber.org/zap"

func CheckErr(e error) {
	if e != nil {
		zap.L().WithOptions(zap.AddCallerSkip(1)).Error(e.Error())
	}
}

func CheckErrWarn(e error) {
	if e != nil {
		zap.L().WithOptions(zap.AddCallerSkip(1)).Warn(e.Error())
	}
}

func CheckErrInfo(e error) {
	if e != nil {
		zap.L().WithOptions(zap.AddCallerSkip(1)).Info(e.Error())
	}
}

func CheckErrFatal(e error) {
	if e != nil {
		zap.L().WithOptions(zap.AddCallerSkip(1)).Fatal(e.Error())
	}
}
