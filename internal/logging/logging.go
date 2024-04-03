package logging

import "go.uber.org/zap"

func New() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()
}
