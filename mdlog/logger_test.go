package mdlog_test

import (
	"context"
	"github.com/chaseisabelle/md/mdlog"
)

type TestLogger struct {
	FatalFunc mdlog.ErrFunc
	ErrorFunc mdlog.ErrFunc
	WarnFunc  mdlog.MsgFunc
	InfoFunc  mdlog.MsgFunc
	DebugFunc mdlog.MsgFunc
}

func (t *TestLogger) Fatal(ctx context.Context, err error, md map[string]any) {
	t.FatalFunc(ctx, err, md)
}

func (t *TestLogger) Error(ctx context.Context, err error, md map[string]any) {
	t.ErrorFunc(ctx, err, md)
}

func (t *TestLogger) Warn(ctx context.Context, msg string, md map[string]any) {
	t.WarnFunc(ctx, msg, md)
}

func (t *TestLogger) Info(ctx context.Context, msg string, md map[string]any) {
	t.InfoFunc(ctx, msg, md)
}

func (t *TestLogger) Debug(ctx context.Context, msg string, md map[string]any) {
	t.DebugFunc(ctx, msg, md)
}
