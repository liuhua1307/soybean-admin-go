package log

import (
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	"golang.org/x/exp/slog"
)

var _ Logger = &SLogLogger{}

type SLogLogger struct {
	logger *slog.Logger
}

const (
	levelTrace = slog.Level(-8)
	levelFatal = slog.Level(12)
	levelPanic = slog.Level(16)
)

var (
	levelNames = map[slog.Leveler]string{
		levelTrace: "TRACE",
		levelFatal: "FATAL",
		levelPanic: "PANIC",
	}
	replaceAttr = func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.LevelKey {
			level := a.Value.Any().(slog.Level)
			levelLabel, exists := levelNames[level]
			if !exists {
				levelLabel = level.String()
			}

			a.Value = slog.StringValue(levelLabel)
		}

		return a
	}
)

func NewSlogLogger(path string) *SLogLogger {
	var handler = getDevHandler(path)
	if os.Getenv("APP_ENV") == "prod" {
		handler = getProdHandler(path)
	}
	slogLogger := slog.New(handler)
	return &SLogLogger{
		logger: slogLogger,
	}
}

func getFileWriter(path string) io.Writer {
	if len(path) == 0 {
		return os.Stdout
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    5,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
	return io.MultiWriter(os.Stdout, lumberJackLogger)
}

func getProdOpts() *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelInfo,
		ReplaceAttr: replaceAttr,
	}
}

func getDevOpts() *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: replaceAttr,
	}
}

func getProdHandler(path string) slog.Handler {
	return slog.NewJSONHandler(getFileWriter(path), getProdOpts())
}

func getDevHandler(path string) slog.Handler {
	return slog.NewTextHandler(getFileWriter(path), getDevOpts())
}

func (s *SLogLogger) Debug(msg string, fields ...Field) {
	s.logger.Debug(msg, s.toSlogAttrs(fields)...)
}

func (s *SLogLogger) Info(msg string, fields ...Field) {
	s.logger.Info(msg, s.toSlogAttrs(fields)...)
}

func (s *SLogLogger) Warn(msg string, fields ...Field) {
	s.logger.Warn(msg, s.toSlogAttrs(fields)...)
}

func (s *SLogLogger) Error(msg string, fields ...Field) {
	s.logger.Error(msg, s.toSlogAttrs(fields)...)
}

func (s *SLogLogger) Fatal(msg string, fields ...Field) {
	s.logger.Log(nil, levelFatal, msg, s.toSlogAttrs(fields)...)
	os.Exit(1)
}

func (s *SLogLogger) Panic(msg string, fields ...Field) {
	s.logger.Log(nil, levelPanic, msg, s.toSlogAttrs(fields)...)
	panic(msg)
}

func (s *SLogLogger) Sync() error {
	return nil
}

func (s *SLogLogger) toSlogAttrs(fields []Field) []any {
	slogAttrs := make([]any, 0, len(fields))
	for _, f := range fields {
		slogAttrs = append(slogAttrs, slog.Any(f.Key, f.Value))
	}
	return slogAttrs
}
