package base

import "log"

// Logger 定义了日志接口，便于调用方自定义日志输出，兼容标准库 log.Logger
type Logger interface {
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

// defaultLogger 是默认的日志实现，直接桥接到标准库的 log
type defaultLogger struct{}

func (d *defaultLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (d *defaultLogger) Println(v ...interface{}) {
	log.Println(v...)
}

var currentLogger Logger = &defaultLogger{}

// SetLogger 设置全局 of Logger 实例，支持调用者注入自定义 Logger
func SetLogger(logger Logger) {
	if logger != nil {
		currentLogger = logger
	}
}

// GetLogger 获取当前全局 of Logger 实例
func GetLogger() Logger {
	return currentLogger
}
