package mdzap

import (
	"context"
	"github.com/chaseisabelle/md/mdlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Zap struct {
	logger *zap.Logger
}

func New(cfg mdlog.Config) (*Zap, error) {
	cll := cfg.Level

	ile := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		var lev mdlog.Level

		switch lvl {
		case zapcore.FatalLevel:
			fallthrough
		case zapcore.ErrorLevel:
			return false
		case zapcore.WarnLevel:
			lev = mdlog.Warn
		case zapcore.InfoLevel:
			lev = mdlog.Info
		case zapcore.DebugLevel:
			lev = mdlog.Debug
		default:
			lev = mdlog.Info
		}

		return lev >= cll
	})

	ele := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		var lev mdlog.Level

		switch lvl {
		case zapcore.FatalLevel:
			lev = mdlog.Fatal
		case zapcore.ErrorLevel:
			lev = mdlog.Error
		case zapcore.WarnLevel:
			fallthrough
		case zapcore.InfoLevel:
			fallthrough
		case zapcore.DebugLevel:
			fallthrough
		default:
			return true
		}

		return lev >= cll
	})

	sos := zapcore.Lock(os.Stdout)
	ses := zapcore.Lock(os.Stderr)
	zoc := zap.NewProductionEncoderConfig()
	zec := zap.NewProductionEncoderConfig()
	zoe := zapcore.NewJSONEncoder(zoc)
	zee := zapcore.NewJSONEncoder(zec)
	ozc := zapcore.NewCore(zoe, sos, ile)
	ezc := zapcore.NewCore(zee, ses, ele)
	tee := zapcore.NewTee(ozc, ezc)
	lgr := zap.New(tee)

	return &Zap{
		logger: lgr,
	}, nil
}

func (z *Zap) Fatal(ctx context.Context, err error, md map[string]any) {
	z.logger.Fatal(err.Error(), metadata(md))
}

func (z *Zap) Error(ctx context.Context, err error, md map[string]any) {
	z.logger.Error(err.Error(), metadata(md))
}

func (z *Zap) Warn(ctx context.Context, msg string, md map[string]any) {
	z.logger.Warn(msg, metadata(md))
}

func (z *Zap) Info(ctx context.Context, msg string, md map[string]any) {
	z.logger.Info(msg, metadata(md))
}

func (z *Zap) Debug(ctx context.Context, msg string, md map[string]any) {
	z.logger.Debug(msg, metadata(md))
}

func metadata(md map[string]any) zapcore.Field {
	if md == nil {
		return zap.Field{}
	}

	fds := make([]zapcore.Field, len(md))
	ind := 0

	for k, v := range md {
		fds[ind] = zap.Any(k, v)

		ind++
	}

	return zap.Any("metadata", fds)
}
