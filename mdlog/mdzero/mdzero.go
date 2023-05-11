package mdzero

import (
	"context"
	"github.com/chaseisabelle/md/mdlog"
	"github.com/chaseisabelle/mderr"
	"github.com/rs/zerolog"
	"os"
)

type Zero struct {
	stdout zerolog.Logger
	stderr zerolog.Logger
}

func New(cfg mdlog.Config) (*Zero, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	sol := zerolog.New(os.Stdout).With().Timestamp().Logger()
	sel := zerolog.New(os.Stderr).With().Timestamp().Logger()
	lvl := cfg.Level

	var zll zerolog.Level

	switch lvl {
	case mdlog.Fatal:
		zll = zerolog.FatalLevel
	case mdlog.Error:
		zll = zerolog.ErrorLevel
	case mdlog.Warn:
		zll = zerolog.WarnLevel
	case mdlog.Info:
		zll = zerolog.InfoLevel
	case mdlog.Debug:
		zll = zerolog.DebugLevel
	default:
		return nil, mderr.New("invalid log level", map[string]any{
			"level": lvl,
		})
	}

	sol = sol.Level(zll)
	sel = sel.Level(zll)

	return &Zero{
		stdout: sol,
		stderr: sel,
	}, nil
}

func (z *Zero) Fatal(ctx context.Context, err error, md map[string]any) {
	z.stderr.Fatal().Err(err).Fields(metadata(md)).Msg("")
}

func (z *Zero) Error(ctx context.Context, err error, md map[string]any) {
	z.stderr.Error().Err(err).Fields(metadata(md)).Msg("")
}

func (z *Zero) Warn(ctx context.Context, msg string, md map[string]any) {
	z.stdout.Warn().Fields(metadata(md)).Msg(msg)
}

func (z *Zero) Info(ctx context.Context, msg string, md map[string]any) {
	z.stdout.Info().Fields(metadata(md)).Msg(msg)
}

func (z *Zero) Debug(ctx context.Context, msg string, md map[string]any) {
	z.stdout.Debug().Fields(metadata(md)).Msg(msg)
}

func metadata(md map[string]any) map[string]any {
	return map[string]any{
		"metadata": md,
	}
}
