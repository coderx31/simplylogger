package simplylogger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger struct {
	stage      string
	baseLogger *zap.Logger
}

func NewLogger(env, stage string) *Logger {
	var encoderCfg zapcore.EncoderConfig
	var dev bool
	if env == "dev" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
		dev = true
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       dev,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	logger := zap.Must(config.Build())
	fields := []zap.Field{
		zap.String("prefix", stage),
	}

	return &Logger{
		stage:      stage,
		baseLogger: logger.With(fields...),
	}
}

func (l *Logger) Debug(msg string, params ...interface{}) {
	l.baseLogger.Debug(msg, zap.Any("params", params))
}

func (l *Logger) DebugContext(ctx context.Context, msg string, params ...interface{}) {
	var traceID string
	var userID string
	if ctx.Value(ContextTraceID) != nil {
		traceID = ctx.Value(ContextTraceID).(string)
	}
	if ctx.Value(ContextUserID) != nil {
		userID = ctx.Value(ContextUserID).(string)
	}
	l.baseLogger.Debug(msg,
		zap.String(ContextTraceID, traceID),
		zap.String(ContextTraceID, userID),
		zap.Any("params", params))
}

func (l *Logger) Info(msg string, params ...interface{}) {
	l.baseLogger.Info(msg, zap.Any("params", params))
}

func (l *Logger) InfoContext(ctx context.Context, msg string, params ...interface{}) {
	var traceID string
	var userID string
	if ctx.Value(ContextTraceID) != nil {
		traceID = ctx.Value(ContextTraceID).(string)
	}
	if ctx.Value(ContextUserID) != nil {
		userID = ctx.Value(ContextUserID).(string)
	}
	l.baseLogger.Info(msg,
		zap.String(ContextTraceID, traceID),
		zap.String(ContextUserID, userID),
		zap.Any("params", params))
}

func (l *Logger) Warn(msg string, params ...interface{}) {
	l.baseLogger.Warn(msg, zap.Any("params", params))
}

func (l *Logger) WarnContext(ctx context.Context, msg string, params ...interface{}) {
	var traceID string
	var userID string
	if ctx.Value(ContextTraceID) != nil {
		traceID = ctx.Value(ContextTraceID).(string)
	}
	if ctx.Value(ContextUserID) != nil {
		userID = ctx.Value(ContextUserID).(string)
	}
	l.baseLogger.Warn(msg,
		zap.String(ContextTraceID, traceID),
		zap.String(ContextUserID, userID),
		zap.Any("params", params))
}

func (l *Logger) Error(msg string, params ...interface{}) {
	l.baseLogger.Error(msg, zap.Any("params", params))
}

func (l *Logger) ErrorContext(ctx context.Context, msg string, params ...interface{}) {
	var traceID string
	var userID string
	if ctx.Value(ContextTraceID) != nil {
		traceID = ctx.Value(ContextTraceID).(string)
	}
	if ctx.Value(ContextUserID) != nil {
		userID = ctx.Value(ContextUserID).(string)
	}
	l.baseLogger.Error(msg,
		zap.String(ContextTraceID, traceID),
		zap.String(ContextUserID, userID),
		zap.Any("params", params))
}

func (l *Logger) Fatal(msg string, params ...interface{}) {
	l.baseLogger.Fatal(msg, zap.Any("params", params))
}

func (l *Logger) FatalContext(ctx context.Context, msg string, params ...interface{}) {
	var traceID string
	var userID string
	if ctx.Value(ContextTraceID) != nil {
		traceID = ctx.Value(ContextTraceID).(string)
	}
	if ctx.Value(ContextUserID) != nil {
		userID = ctx.Value(ContextUserID).(string)
	}
	l.baseLogger.Fatal(msg,
		zap.String(ContextTraceID, traceID),
		zap.String(ContextUserID, userID),
		zap.Any("params", params))
}
