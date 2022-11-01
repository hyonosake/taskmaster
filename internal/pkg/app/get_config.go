package app

import (
	"os"

	"go.uber.org/zap/zapcore"
)

type LogLevel uint8

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
)

type Config struct {
	ConfigPath string
	LogFile    *os.File
	LogLevel   zapcore.Level
}

func (c *Config) GetConfigPath() string {
	return c.ConfigPath
}

func (c *Config) GetLogFile() *os.File {
	return c.LogFile
}

func (c *Config) GetLogLevel() zapcore.Level {
	switch c.LogLevel {
	case DEBUG:
		return zapcore.DebugLevel
	case INFO:
		return zapcore.InfoLevel
	case WARNING:
		return zapcore.WarnLevel
	case ERROR:
		return zapcore.ErrorLevel
	}
	return zapcore.ErrorLevel
}
