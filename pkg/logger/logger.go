package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var logger *zap.Logger

func init() {
	conf := zap.NewProductionConfig()
	conf.Level.SetLevel(zapcore.DebugLevel)

	zapLgr, err := conf.Build()
	if err != nil {
		log.Fatal("Error while creating logger")
	}
	logger = zapLgr.With(zap.String("service", "library-api"))
}

func Logger() *zap.Logger {
	return logger
}
