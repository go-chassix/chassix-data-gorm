package gormx

import (
	"c5x.io/logx"
	"gorm.io/gorm/logger"
)

type Logger struct {
	entry *logx.Entry
}

func NewLogger() *Logger {
	return new(Logger)
}

func (l *Logger) LogModel(level logger.LogLevel) {

}
