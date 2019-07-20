package logger

import "log"

type Logger struct {
	Enable bool
}

func (m *Logger) Debug(args ...interface{}) {
	if m.Enable {
		m.Print(args...)
	}
}

func (m *Logger) Print(args ...interface{}) {
	log.Print(args...)
}
