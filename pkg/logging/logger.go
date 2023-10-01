package logging

import (
	"context"
)

type AppLogger interface {
	Debugf(msg string, args ...any)
	Infof(msg string, args ...any)
	Errorf(msg string, args ...any)
	Fatalf(msg string, args ...any)
	Debug(msg string)
	Info(msg string)
	Error(msg string)
	Fatal(msg string)
	WithFields(map[string]string) AppLogger
	WithContext(ctx context.Context) AppLogger
}
