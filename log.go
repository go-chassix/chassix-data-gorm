package gormx

import (
	"c5x.io/logx"
	"context"
	"gorm.io/gorm/logger"
	"time"
)

type Logger struct {
	entry *logx.Entry
	logger.Config
}

func NewLogger(logCfg *LoggerConfig) *Logger {
	log := new(Logger)
	log.entry = logx.New().Category("gorm")
	log.Config = logger.Config{
		SlowThreshold: logCfg.SlowThreshold,
		Colorful:      logCfg.Colorful,
		LogLevel:      logCfg.Level,
	}
	return log
}

func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {

	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func (l *Logger) Info(ctx context.Context, str string, args ...interface{}) {
	l.entry.Info(str, args)
}

func (l *Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	l.entry.Info(str, args)
}
func (l *Logger) Error(ctx context.Context, str string, args ...interface{}) {
	l.entry.Error(str, args)
}
func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	l.entry.Trace("begin:%s, error: %v", begin.String(), err)
}
