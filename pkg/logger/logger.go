package logger

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type traceIDKey string

const traceID traceIDKey = "traceID"

const (
	productionEnv  = "PROD"
	developmentEnv = "DEV"
)

type Logger interface {
	Info(ctx context.Context, args ...interface{})
	Infof(ctx context.Context, template string, args ...interface{})
	Debug(ctx context.Context, args ...interface{})
	Debugf(ctx context.Context, template string, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Errorf(ctx context.Context, template string, args ...interface{})
	Fatal(ctx context.Context, args ...interface{})
	Fatalf(ctx context.Context, template string, args ...interface{})
	SetTraceID(ctx context.Context) context.Context
}

type logger struct {
	zap *zap.SugaredLogger
}

func New(level string, environment string) (Logger, error) {
	var config zap.Config

	if environment == productionEnv {
		config = zap.NewProductionConfig()
	}
	if environment == developmentEnv {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	config.DisableCaller = true
	config.DisableStacktrace = true
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		return nil, err
	}
	config.Level.SetLevel(zapLevel)

	zapLogger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	return &logger{
		zap: zapLogger.Sugar(),
	}, nil

}

func (l *logger) SetTraceID(ctx context.Context) context.Context {
	traceId := uuid.NewString()
	return context.WithValue(ctx, traceID, traceId)
}

func (l *logger) GetTraceID(ctx context.Context) string {
	v := ctx.Value(traceID)
	if v == nil {
		return l.SetTraceID(ctx).Value(traceID).(string)
	}
	return v.(string)
}

func (l *logger) withTraceID(ctx context.Context) zap.Field {
	return zap.String("traceID", l.GetTraceID(ctx))
}

func (l *logger) Info(ctx context.Context, args ...interface{}) {
	l.zap.With(l.withTraceID(ctx)).Info(args)
}

func (l *logger) Infof(ctx context.Context, template string, args ...interface{}) {
	l.zap.With(l.withTraceID(ctx)).Infof(template, args)
}
func (l *logger) Debug(ctx context.Context, args ...interface{}) {
	l.zap.With(l.withTraceID(ctx)).Debug(args)
}
func (l *logger) Debugf(ctx context.Context, template string, args ...interface{}) {
	l.zap.With(l.withTraceID(ctx)).Debugf(template, args)
}
func (l *logger) Error(ctx context.Context, args ...interface{}) {
	l.zap.With(l.withTraceID(ctx)).Error(args)
}
func (l *logger) Errorf(ctx context.Context, template string, args ...interface{}) {
	l.zap.With(l.withTraceID(ctx)).Errorf(template, args)
}
func (l *logger) Fatal(ctx context.Context, args ...interface{}) {
	l.zap.With(l.withTraceID(ctx)).Fatal(args)
}
func (l *logger) Fatalf(ctx context.Context, template string, args ...interface{}) {
	l.zap.With(l.withTraceID(ctx)).Fatalf(template, args)
}