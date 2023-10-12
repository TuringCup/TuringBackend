package logger

import "go.uber.org/zap"

var Logger *zap.Logger

func LoggerInit() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		panic("logger init failed")
	}
}
