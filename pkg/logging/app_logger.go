package logging

import (
	"context"
	"fmt"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"sync"
)

var Logger ApplicationLogger
var once sync.Once

type ApplicationLogger struct {
	logger *ZapLogger
}

func InitLogger(lvl, serviceName, environment string) error {
	level, err := parseLevel(lvl)
	Logger = getLogger(level, serviceName, environment)
	return err
}

func getLogger(level zapcore.Level, serviceName string, environment string) ApplicationLogger {
	once.Do(func() {
		encoderConfig := ecszap.EncoderConfig{
			EncodeName:     zap.NewProductionEncoderConfig().EncodeName,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   ecszap.FullCallerEncoder,
		}
		core := ecszap.NewCore(encoderConfig, os.Stdout, level)
		l := zap.New(core, zap.AddCaller())
		l = l.With(zap.String("app", serviceName)).With(zap.String("env", environment))

		zapLogger := NewZapLogger(l)

		Logger = ApplicationLogger{
			logger: zapLogger,
		}
	})

	return Logger
}

func (l *ApplicationLogger) Debugf(msg string, args ...any) {
	l.logger.Debugf(msg, args...)
}
func (l *ApplicationLogger) Infof(msg string, args ...any) {
	l.logger.Infof(msg, args...)
}
func (l *ApplicationLogger) Errorf(msg string, args ...any) {
	l.logger.Errorf(msg, args...)
}
func (l *ApplicationLogger) Fatalf(msg string, args ...any) {
	l.logger.Fatalf(msg, args...)
}
func (l *ApplicationLogger) Debug(msg string) {
	l.logger.Debug(msg)
}
func (l *ApplicationLogger) Info(msg string) {
	l.logger.Info(msg)
}
func (l *ApplicationLogger) Error(msg string) {
	l.logger.Error(msg)
}
func (l *ApplicationLogger) Fatal(msg string) {
	l.logger.Fatal(msg)
}
func (l *ApplicationLogger) WithFields(fields map[string]string) AppLogger {
	return l.logger.WithFields(fields)
}
func (l *ApplicationLogger) WithContext(ctx context.Context) AppLogger {
	return l.logger.WithContext(ctx)
}

func parseLevel(lvl string) (zapcore.Level, error) {
	switch strings.ToLower(lvl) {
	case "debug":
		return zap.DebugLevel, nil
	case "info":
		return zap.InfoLevel, nil
	case "warn":
		return zap.WarnLevel, nil
	case "error":
		return zap.ErrorLevel, nil
	}
	return zap.InfoLevel, fmt.Errorf("invalid log level <%v>", lvl)
}
