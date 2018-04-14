package logger

import (
	"log"
	"os"
	"strings"
)

const (
	stateDebug = "DEBUG"
	stateInfo = "INFO"
	stateWarnning = "WARNING"
	stateError = "ERROR"
	stateFatal = "FATAL"
)

var (
	instance *Logger
	logger   *log.Logger
)

type Logger struct {
	level int
}

func New() *Logger {
	if instance == nil {
		instance = &Logger{DefaultLoggerLevel}
		if logger == nil {
			logger = log.New(os.Stdout, "", log.LstdFlags)
		}
	}
	return instance
}

func createMessage(v ...interface{}) string {
	s := make([]string, len(v))
	for i, a := range v {
		s[i] = a.(string)
	}
	return strings.Join(s, " ")
}

func (l *Logger) Print(v ...interface{}) {
	logger.Print(v)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	logger.Printf(format, v)
}

func (l *Logger) Println(v ...interface{}) {
	logger.Println(v)
}

func (l *Logger) Debug(v ...interface{}) {
	if l.level <= LevelDebug {
		logger.Printf("%s: %s", stateDebug, createMessage(v))
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level <= LevelDebug {
		logger.Printf(format, v)
	}
}

func (l *Logger) Debugln(v ...interface{}) {
	if l.level <= LevelDebug {
		logger.Println(v)
	}
}
