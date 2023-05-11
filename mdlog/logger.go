package mdlog

import (
	"context"
)

// Logger is the interface for all logger implementations
type Logger interface {
	Fatal(context.Context, error, map[string]any)
	Error(context.Context, error, map[string]any)
	Warn(context.Context, string, map[string]any)
	Info(context.Context, string, map[string]any)
	Debug(context.Context, string, map[string]any)
}

// ErrFunc is a func that handles an error entry
type ErrFunc func(context.Context, error, map[string]any)

// MsgFunc is a func that handles a message entry
type MsgFunc func(context.Context, string, map[string]any)
