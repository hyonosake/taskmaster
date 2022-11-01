package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewSugaredLogger(lvl zapcore.Level, file *os.File) (*zap.SugaredLogger, error) {

	pe := zap.NewProductionEncoderConfig()

	fileEncoder := zapcore.NewJSONEncoder(pe)

	pe.EncodeTime = zapcore.RFC3339TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	var core zapcore.Core
	if file == nil {
		core = zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), lvl)
	} else {
		core = zapcore.NewTee(
			zapcore.NewCore(fileEncoder, zapcore.AddSync(file), lvl),
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), lvl),
		)
	}

	l := zap.New(core)

	return l.Sugar(), nil
}
