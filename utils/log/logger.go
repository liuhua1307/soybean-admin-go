// Package log，定义了一个日志通用接口。Debug Info Warn Error Fatal Panic Sync 等方法，以及 Field 结构体。
// 通过定义 Logger 日志接口和定义日志数据结构，能够将日志模块和项目解耦，当日志模块选用发生改变时，只需要修改日志模块，再将 /internal/pkg/log/log.go 中的 log 变量修改为新的日志模块即可。
// 对项目中的日志调用并不会影响。
package log

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	Panic(msg string, fields ...Field)
	Sync() error
}

type Field struct {
	Key   string
	Value interface{}
}
