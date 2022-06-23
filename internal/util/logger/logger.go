package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

const (
	DEFAULT   Level = 0
	DEBUG     Level = 100
	INFO      Level = 200
	NOTICE    Level = 300
	WARNING   Level = 400
	ERROR     Level = 500
	CRITICAL  Level = 600
	ALERT     Level = 700
	EMERGENCY Level = 800
)

type (
	Level uint16

	Logger struct {
		name        string
		level       Level
		output      io.Writer
		errorOutput io.Writer
		levels      []string
		appendices  []Appendix
		mutex       sync.Mutex
	}
)

func New(name string, level Level, options ...Option) *Logger {
	logger := &Logger{
		name:        name,
		level:       level,
		output:      os.Stdout,
		errorOutput: os.Stderr,
		levels: []string{
			DEFAULT:   "DEFAULT",
			DEBUG:     "DEBUG",
			INFO:      "INFO",
			NOTICE:    "NOTICE",
			WARNING:   "WARNING",
			ERROR:     "ERROR",
			CRITICAL:  "CRITICAL",
			ALERT:     "ALERT",
			EMERGENCY: "EMERGENCY",
		},
	}
	for _, option := range options {
		option.apply(logger)
	}
	return logger
}

func (logger *Logger) Default(message string, data map[string]interface{}) {
	logger.log(context.Background(), DEFAULT, message, data)
}

func (logger *Logger) Debug(message string, data map[string]interface{}) {
	logger.log(context.Background(), DEBUG, message, data)
}

func (logger *Logger) Info(message string, data map[string]interface{}) {
	logger.log(context.Background(), INFO, message, data)
}

func (logger *Logger) Notice(message string, data map[string]interface{}) {
	logger.log(context.Background(), NOTICE, message, data)
}

func (logger *Logger) Warning(message string, data map[string]interface{}) {
	logger.log(context.Background(), WARNING, message, data)
}

func (logger *Logger) Error(message string, data map[string]interface{}) {
	logger.log(context.Background(), ERROR, message, data)
}

func (logger *Logger) Critical(message string, data map[string]interface{}) {
	logger.log(context.Background(), CRITICAL, message, data)
}

func (logger *Logger) Alert(message string, data map[string]interface{}) {
	logger.log(context.Background(), ALERT, message, data)
}

func (logger *Logger) Emergency(message string, data map[string]interface{}) {
	logger.log(context.Background(), EMERGENCY, message, data)
}

func (logger *Logger) DefaultContext(ctx context.Context, message string, data map[string]interface{}) {
	logger.log(ctx, DEFAULT, message, data)
}

func (logger *Logger) DebugContext(ctx context.Context, message string, data map[string]interface{}) {
	logger.log(ctx, DEBUG, message, data)
}

func (logger *Logger) InfoContext(ctx context.Context, message string, data map[string]interface{}) {
	logger.log(ctx, INFO, message, data)
}

func (logger *Logger) NoticeContext(ctx context.Context, message string, data map[string]interface{}) {
	logger.log(ctx, NOTICE, message, data)
}

func (logger *Logger) WarningContext(ctx context.Context, message string, data map[string]interface{}) {
	logger.log(ctx, WARNING, message, data)
}

func (logger *Logger) ErrorContext(ctx context.Context, message string, data map[string]interface{}) {
	logger.log(ctx, ERROR, message, data)
}

func (logger *Logger) CriticalContext(ctx context.Context, message string, data map[string]interface{}) {
	logger.log(ctx, CRITICAL, message, data)
}

func (logger *Logger) AlertContext(ctx context.Context, message string, data map[string]interface{}) {
	logger.log(ctx, ALERT, message, data)
}

func (logger *Logger) EmergencyContext(ctx context.Context, message string, data map[string]interface{}) {
	logger.log(ctx, EMERGENCY, message, data)
}

func (logger *Logger) log(ctx context.Context, level Level, message string, data map[string]interface{}) {
	if level < logger.level {
		return
	}
	record := map[string]interface{}{
		"name":      logger.name,
		"message":   message,
		"context":   data,
		"timestamp": time.Now().Format(time.RFC3339Nano),
		"severity":  logger.levels[level],
	}
	for _, appendix := range logger.appendices {
		record = appendix.run(ctx, level, record)
	}
	b, err := json.Marshal(record)
	if err != nil {
		return
	}
	b = append(b, '\n')
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	if level < WARNING {
		_, _ = fmt.Fprint(logger.output, string(b))
	} else {
		_, _ = fmt.Fprint(logger.errorOutput, string(b))
	}
}
