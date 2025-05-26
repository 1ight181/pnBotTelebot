package logruslogger

import (
	"context"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	Entry *logrus.Entry
}

func (l *LogrusLogger) Debug(msg string) {
	l.Entry.Debug(msg)
}
func (l *LogrusLogger) Debugf(format string, args ...interface{}) {
	l.Entry.Debugf(format, args...)
}

func (l *LogrusLogger) Info(msg string) {
	l.Entry.Info(msg)
}
func (l *LogrusLogger) Infof(format string, args ...interface{}) {
	l.Entry.Infof(format, args...)
}

func (l *LogrusLogger) Warn(msg string) {
	l.Entry.Warn(msg)
}
func (l *LogrusLogger) Warnf(format string, args ...interface{}) {
	l.Entry.Warnf(format, args...)
}

func (l *LogrusLogger) Error(msg string) {
	l.Entry.Error(msg)
}
func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	l.Entry.Errorf(format, args...)
}

func (l *LogrusLogger) Fatal(msg string) {
	l.Entry.Fatal(msg)
}
func (l *LogrusLogger) Fatalf(format string, args ...interface{}) {
	l.Entry.Fatalf(format, args...)
}

func (l *LogrusLogger) Panic(msg string) {
	l.Entry.Panic(msg)
}
func (l *LogrusLogger) Panicf(format string, args ...interface{}) {
	l.Entry.Panicf(format, args...)
}

func (l *LogrusLogger) SetContext(ctx context.Context) {
	l.Entry = l.Entry.WithContext(ctx)
}
