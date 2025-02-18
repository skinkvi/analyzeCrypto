package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout"}

	var err error
	Logger, err = config.Build()
	if err != nil {
		panic(err)
	}

	defer Logger.Sync()
}
